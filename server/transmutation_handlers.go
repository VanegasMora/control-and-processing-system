package server

import (
	"backend-avanzada/api"
	"backend-avanzada/auth"
	"backend-avanzada/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (s *Server) HandleTransmutations(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetAllTransmutations(w, r)
	case http.MethodPost:
		s.handleCreateTransmutation(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) HandleTransmutationByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetTransmutationByID(w, r)
	case http.MethodPut:
		s.handleUpdateTransmutation(w, r)
	case http.MethodDelete:
		s.handleDeleteTransmutation(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleGetAllTransmutations(w http.ResponseWriter, r *http.Request) {
	userID, _ := auth.GetUserID(r)
	userRole, _ := auth.GetUserRole(r)

	var transmutations []*models.Transmutation
	var err error

	if userRole == "supervisor" {
		transmutations, err = s.TransmutationRepository.FindAll()
	} else {
		transmutations, err = s.TransmutationRepository.FindByAlchemistID(userID)
	}

	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	response := make([]api.TransmutationResponse, len(transmutations))
	for i, t := range transmutations {
		response[i] = s.transmutationToResponse(t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleGetTransmutationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	transmutation, err := s.TransmutationRepository.FindById(uint(id))
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	if transmutation == nil {
		http.Error(w, "Transmutation not found", http.StatusNotFound)
		return
	}

	response := s.transmutationToResponse(transmutation)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleCreateTransmutation(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.GetUserID(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req api.TransmutationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Calcular costo
	cost := 0.0
	for _, input := range req.InputMaterials {
		material, err := s.MaterialRepository.FindById(input.MaterialID)
		if err != nil {
			s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
			return
		}
		if material == nil {
			http.Error(w, fmt.Sprintf("Material %d not found", input.MaterialID), http.StatusNotFound)
			return
		}
		cost += material.Price * input.Quantity
	}

	transmutation := &models.Transmutation{
		AlchemistID: userID,
		Status:      models.TransmutationStatusPending,
		Description: req.Description,
		Cost:        cost,
	}

	transmutation, err := s.TransmutationRepository.Create(transmutation)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	// Crear materiales de entrada
	for _, input := range req.InputMaterials {
		tm := &models.TransmutationMaterial{
			TransmutationID: transmutation.ID,
			MaterialID:      input.MaterialID,
			Quantity:        input.Quantity,
			IsInput:         true,
		}
		if err := s.DB.Create(tm).Error; err != nil {
			s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
			return
		}
	}

	// Crear materiales de salida si existen
	for _, output := range req.OutputMaterials {
		tm := &models.TransmutationMaterial{
			TransmutationID: transmutation.ID,
			MaterialID:      output.MaterialID,
			Quantity:        output.Quantity,
			IsInput:         false,
		}
		if err := s.DB.Create(tm).Error; err != nil {
			s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
			return
		}
	}

	// Enviar a cola de tareas para procesamiento asíncrono
	s.taskQueue.Enqueue("transmutation_created", map[string]interface{}{
		"transmutation_id": transmutation.ID,
		"alchemist_id":     userID,
		"cost":             cost,
	})

	// Recargar con relaciones
	transmutation, _ = s.TransmutationRepository.FindById(transmutation.ID)
	response := s.transmutationToResponse(transmutation)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (s *Server) HandleUpdateTransmutationStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	transmutation, err := s.TransmutationRepository.FindById(uint(id))
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	if transmutation == nil {
		http.Error(w, "Transmutation not found", http.StatusNotFound)
		return
	}

	var req api.TransmutationStatusUpdate
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID, _ := auth.GetUserID(r)
	now := time.Now()

	transmutation.Status = models.TransmutationStatus(req.Status)
	transmutation.Result = req.Result

	if req.Status == string(models.TransmutationStatusApproved) {
		transmutation.ApprovedAt = &now
		transmutation.SupervisorID = &userID
	} else if req.Status == string(models.TransmutationStatusCompleted) {
		transmutation.CompletedAt = &now
		// Procesar transmutación: actualizar stocks
		s.processTransmutation(transmutation)
	}

	transmutation, err = s.TransmutationRepository.Save(transmutation)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	// Notificar cambio de estado
	s.NotifyWebSocket("transmutation_status_changed", map[string]interface{}{
		"transmutation_id": transmutation.ID,
		"status":           transmutation.Status,
		"alchemist_id":     transmutation.AlchemistID,
	})

	response := s.transmutationToResponse(transmutation)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) processTransmutation(transmutation *models.Transmutation) {
	// Recargar con materiales
	t, _ := s.TransmutationRepository.FindById(transmutation.ID)
	if t == nil {
		return
	}

	// Reducir stock de materiales de entrada
	for _, input := range t.InputMaterials {
		s.MaterialRepository.UpdateStock(input.MaterialID, -input.Quantity)
	}

	// Aumentar stock de materiales de salida
	for _, output := range t.OutputMaterials {
		s.MaterialRepository.UpdateStock(output.MaterialID, output.Quantity)
	}
}

func (s *Server) handleUpdateTransmutation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	transmutation, err := s.TransmutationRepository.FindById(uint(id))
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	if transmutation == nil {
		http.Error(w, "Transmutation not found", http.StatusNotFound)
		return
	}

	var req api.TransmutationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	transmutation.Description = req.Description

	transmutation, err = s.TransmutationRepository.Save(transmutation)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	response := s.transmutationToResponse(transmutation)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleDeleteTransmutation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	transmutation, err := s.TransmutationRepository.FindById(uint(id))
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	if transmutation == nil {
		http.Error(w, "Transmutation not found", http.StatusNotFound)
		return
	}

	if err := s.TransmutationRepository.Delete(transmutation); err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) transmutationToResponse(t *models.Transmutation) api.TransmutationResponse {
	var alchemist *api.AlchemistResponse
	if t.Alchemist.ID != 0 {
		alchemist = &api.AlchemistResponse{
			ID:        t.Alchemist.ID,
			Name:      t.Alchemist.Name,
			Email:     t.Alchemist.Email,
			Rank:      string(t.Alchemist.Rank),
			Specialty: string(t.Alchemist.Specialty),
			Role:      string(t.Alchemist.Role),
			Certified: t.Alchemist.Certified,
		}
	}

	inputMaterials := make([]api.TransmutationMaterialResponse, len(t.InputMaterials))
	for i, m := range t.InputMaterials {
		inputMaterials[i] = api.TransmutationMaterialResponse{
			ID: m.ID,
			Material: &api.MaterialResponse{
				ID:          m.Material.ID,
				Name:        m.Material.Name,
				Type:        string(m.Material.Type),
				Description: m.Material.Description,
				Stock:       m.Material.Stock,
				Unit:        m.Material.Unit,
				Price:       m.Material.Price,
			},
			Quantity: m.Quantity,
			IsInput:  m.IsInput,
		}
	}

	outputMaterials := make([]api.TransmutationMaterialResponse, len(t.OutputMaterials))
	for i, m := range t.OutputMaterials {
		outputMaterials[i] = api.TransmutationMaterialResponse{
			ID: m.ID,
			Material: &api.MaterialResponse{
				ID:          m.Material.ID,
				Name:        m.Material.Name,
				Type:        string(m.Material.Type),
				Description: m.Material.Description,
				Stock:       m.Material.Stock,
				Unit:        m.Material.Unit,
				Price:       m.Material.Price,
			},
			Quantity: m.Quantity,
			IsInput:  m.IsInput,
		}
	}

	response := api.TransmutationResponse{
		ID:              t.ID,
		AlchemistID:     t.AlchemistID,
		Alchemist:       alchemist,
		Status:          string(t.Status),
		InputMaterials:  inputMaterials,
		OutputMaterials: outputMaterials,
		Description:     t.Description,
		Cost:            t.Cost,
		Result:          t.Result,
		SupervisorID:    t.SupervisorID,
		CreatedAt:       t.CreatedAt.Format(time.RFC3339),
	}

	if t.ApprovedAt != nil {
		approvedAt := t.ApprovedAt.Format(time.RFC3339)
		response.ApprovedAt = &approvedAt
	}
	if t.CompletedAt != nil {
		completedAt := t.CompletedAt.Format(time.RFC3339)
		response.CompletedAt = &completedAt
	}

	return response
}
