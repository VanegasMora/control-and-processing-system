package server

import (
	"backend-avanzada/models"
	"encoding/json"
	"fmt"
	"time"
)

func (s *Server) startAsyncProcessor() {
	ticker := time.NewTicker(1 * time.Minute) // Verificar cada minuto
	defer ticker.Stop()

	// Procesar tareas pendientes inmediatamente
	s.processPendingTasks()

	for {
		select {
		case <-ticker.C:
			s.processPendingTasks()
			s.runDailyChecks()
		}
	}
}

func (s *Server) processPendingTasks() {
	// Procesar transmutaciones pendientes
	pendingTransmutations, err := s.TransmutationRepository.FindByStatus(models.TransmutationStatusPending)
	if err == nil {
		for _, t := range pendingTransmutations {
			s.processTransmutationRequest(t)
		}
	}

	// Procesar misiones pendientes
	pendingMissions, err := s.MissionRepository.FindByStatus(models.MissionStatusPending)
	if err == nil {
		for _, m := range pendingMissions {
			s.processMissionRequest(m)
		}
	}
}

func (s *Server) processTransmutationRequest(transmutation *models.Transmutation) {
	// Verificar disponibilidad de materiales
	transmutation, _ = s.TransmutationRepository.FindById(transmutation.ID)
	if transmutation == nil {
		return
	}

	canProceed := true
	for _, input := range transmutation.InputMaterials {
		material, err := s.MaterialRepository.FindById(input.MaterialID)
		if err != nil || material == nil {
			canProceed = false
			break
		}
		if material.Stock < input.Quantity {
			canProceed = false
			// Crear auditoría
			s.createAudit(models.AuditTypeMaterialUsage, models.AuditSeverityHigh,
				fmt.Sprintf("Insufficient stock for material %s in transmutation %d", material.Name, transmutation.ID),
				&transmutation.AlchemistID,
				map[string]interface{}{
					"transmutation_id": transmutation.ID,
					"material_id":      input.MaterialID,
					"required":         input.Quantity,
					"available":        material.Stock,
				})
			break
		}
	}

	if !canProceed {
		return
	}

	// Simular resultado de transmutación
	result := fmt.Sprintf("Transmutación completada exitosamente. Costo: %.2f", transmutation.Cost)
	transmutation.Result = result

	// Notificar al alquimista
	s.NotifyUser(transmutation.AlchemistID, "transmutation_ready", map[string]interface{}{
		"transmutation_id": transmutation.ID,
		"status":           "ready_for_approval",
	})
}

func (s *Server) processMissionRequest(mission *models.Mission) {
	// Verificar misiones no cerradas por mucho tiempo
	if time.Since(mission.RequestedAt) > 7*24*time.Hour {
		s.createAudit(models.AuditTypeMissionCheck, models.AuditSeverityMedium,
			fmt.Sprintf("Mission %d has been pending for more than 7 days", mission.ID),
			&mission.AlchemistID,
			map[string]interface{}{
				"mission_id":   mission.ID,
				"days_pending": int(time.Since(mission.RequestedAt).Hours() / 24),
			})
	}
}

func (s *Server) runDailyChecks() {
	// Verificar uso excesivo de materiales
	s.checkMaterialUsage()

	// Verificar misiones no cerradas
	s.checkUnclosedMissions()
}

func (s *Server) checkMaterialUsage() {
	// Obtener todas las transmutaciones del último día
	transmutations, err := s.TransmutationRepository.FindAll()
	if err != nil {
		return
	}

	materialUsage := make(map[uint]float64)
	alchemistUsage := make(map[uint]map[uint]float64)

	oneDayAgo := time.Now().Add(-24 * time.Hour)
	for _, t := range transmutations {
		if t.CreatedAt.Before(oneDayAgo) {
			continue
		}

		if alchemistUsage[t.AlchemistID] == nil {
			alchemistUsage[t.AlchemistID] = make(map[uint]float64)
		}

		for _, input := range t.InputMaterials {
			materialUsage[input.MaterialID] += input.Quantity
			alchemistUsage[t.AlchemistID][input.MaterialID] += input.Quantity
		}
	}

	// Verificar uso excesivo por alquimista
	for alchemistID, usage := range alchemistUsage {
		for materialID, quantity := range usage {
			material, _ := s.MaterialRepository.FindById(materialID)
			if material != nil && quantity > material.Stock*0.5 {
				s.createAudit(models.AuditTypeMaterialUsage, models.AuditSeverityHigh,
					fmt.Sprintf("Excessive material usage detected for alchemist %d", alchemistID),
					&alchemistID,
					map[string]interface{}{
						"material_id": materialID,
						"quantity":    quantity,
						"threshold":   material.Stock * 0.5,
					})
			}
		}
	}
}

func (s *Server) checkUnclosedMissions() {
	missions, err := s.MissionRepository.FindAll()
	if err != nil {
		return
	}

	for _, m := range missions {
		if m.Status == models.MissionStatusInProgress {
			// Verificar si la misión lleva más de 30 días en progreso
			if m.ApprovedAt != nil && time.Since(*m.ApprovedAt) > 30*24*time.Hour {
				s.createAudit(models.AuditTypeMissionCheck, models.AuditSeverityMedium,
					fmt.Sprintf("Mission %d has been in progress for more than 30 days", m.ID),
					&m.AlchemistID,
					map[string]interface{}{
						"mission_id":       m.ID,
						"days_in_progress": int(time.Since(*m.ApprovedAt).Hours() / 24),
					})
			}
		}
	}
}

func (s *Server) createAudit(auditType models.AuditType, severity models.AuditSeverity, description string, alchemistID *uint, details map[string]interface{}) {
	detailsJSON, _ := json.Marshal(details)

	audit := &models.Audit{
		Type:        auditType,
		Severity:    severity,
		Description: description,
		AlchemistID: alchemistID,
		Details:     string(detailsJSON),
		Resolved:    false,
	}

	_, err := s.AuditRepository.Create(audit)
	if err != nil {
		s.logger.Error(500, "/audits", err)
	}
}
