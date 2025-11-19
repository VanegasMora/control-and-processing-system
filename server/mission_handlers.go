package server

import (
	"backend-avanzada/api"
	"backend-avanzada/auth"
	"backend-avanzada/models"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (s *Server) HandleMissions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetAllMissions(w, r)
	case http.MethodPost:
		s.handleCreateMission(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) HandleMissionByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetMissionByID(w, r)
	case http.MethodPut:
		s.handleUpdateMission(w, r)
	case http.MethodDelete:
		s.handleDeleteMission(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleGetAllMissions(w http.ResponseWriter, r *http.Request) {
	userID, _ := auth.GetUserID(r)
	userRole, _ := auth.GetUserRole(r)

	var missions []*models.Mission
	var err error

	if userRole == "supervisor" {
		missions, err = s.MissionRepository.FindAll()
	} else {
		missions, err = s.MissionRepository.FindByAlchemistID(userID)
	}

	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	response := make([]api.MissionResponse, len(missions))
	for i, m := range missions {
		var alchemist *api.AlchemistResponse
		if m.Alchemist.ID != 0 {
			alchemist = &api.AlchemistResponse{
				ID:        m.Alchemist.ID,
				Name:      m.Alchemist.Name,
				Email:     m.Alchemist.Email,
				Rank:      string(m.Alchemist.Rank),
				Specialty: string(m.Alchemist.Specialty),
				Role:      string(m.Alchemist.Role),
				Certified: m.Alchemist.Certified,
			}
		}

		response[i] = api.MissionResponse{
			ID:           m.ID,
			Title:        m.Title,
			Description:  m.Description,
			Status:       string(m.Status),
			AlchemistID:  m.AlchemistID,
			Alchemist:    alchemist,
			RequestedAt:  m.RequestedAt.Format(time.RFC3339),
			SupervisorID: m.SupervisorID,
		}

		if m.ApprovedAt != nil {
			approvedAt := m.ApprovedAt.Format(time.RFC3339)
			response[i].ApprovedAt = &approvedAt
		}
		if m.CompletedAt != nil {
			completedAt := m.CompletedAt.Format(time.RFC3339)
			response[i].CompletedAt = &completedAt
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleGetMissionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mission, err := s.MissionRepository.FindById(uint(id))
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	if mission == nil {
		http.Error(w, "Mission not found", http.StatusNotFound)
		return
	}

	var alchemist *api.AlchemistResponse
	if mission.Alchemist.ID != 0 {
		alchemist = &api.AlchemistResponse{
			ID:        mission.Alchemist.ID,
			Name:      mission.Alchemist.Name,
			Email:     mission.Alchemist.Email,
			Rank:      string(mission.Alchemist.Rank),
			Specialty: string(mission.Alchemist.Specialty),
			Role:      string(mission.Alchemist.Role),
			Certified: mission.Alchemist.Certified,
		}
	}

	response := api.MissionResponse{
		ID:           mission.ID,
		Title:        mission.Title,
		Description:  mission.Description,
		Status:       string(mission.Status),
		AlchemistID:  mission.AlchemistID,
		Alchemist:    alchemist,
		RequestedAt:  mission.RequestedAt.Format(time.RFC3339),
		SupervisorID: mission.SupervisorID,
	}

	if mission.ApprovedAt != nil {
		approvedAt := mission.ApprovedAt.Format(time.RFC3339)
		response.ApprovedAt = &approvedAt
	}
	if mission.CompletedAt != nil {
		completedAt := mission.CompletedAt.Format(time.RFC3339)
		response.CompletedAt = &completedAt
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleCreateMission(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.GetUserID(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req api.MissionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	alchemistID := userID
	if req.AlchemistID != 0 {
		alchemistID = req.AlchemistID
	}

	mission := &models.Mission{
		Title:       req.Title,
		Description: req.Description,
		Status:      models.MissionStatusPending,
		AlchemistID: alchemistID,
		RequestedAt: time.Now(),
	}

	mission, err := s.MissionRepository.Create(mission)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	// Enviar a cola de tareas para procesamiento asíncrono
	s.taskQueue.Enqueue("mission_created", map[string]interface{}{
		"mission_id":   mission.ID,
		"alchemist_id": mission.AlchemistID,
	})

	response := api.MissionResponse{
		ID:          mission.ID,
		Title:       mission.Title,
		Description: mission.Description,
		Status:      string(mission.Status),
		AlchemistID: mission.AlchemistID,
		RequestedAt: mission.RequestedAt.Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleUpdateMission(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mission, err := s.MissionRepository.FindById(uint(id))
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	if mission == nil {
		http.Error(w, "Mission not found", http.StatusNotFound)
		return
	}

	var req api.MissionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	mission.Title = req.Title
	mission.Description = req.Description

	mission, err = s.MissionRepository.Save(mission)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	response := api.MissionResponse{
		ID:          mission.ID,
		Title:       mission.Title,
		Description: mission.Description,
		Status:      string(mission.Status),
		AlchemistID: mission.AlchemistID,
		RequestedAt: mission.RequestedAt.Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) HandleUpdateMissionStatus(w http.ResponseWriter, r *http.Request) {
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

	mission, err := s.MissionRepository.FindById(uint(id))
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	if mission == nil {
		http.Error(w, "Mission not found", http.StatusNotFound)
		return
	}

	var req api.MissionStatusUpdate
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID, _ := auth.GetUserID(r)
	now := time.Now()

	mission.Status = models.MissionStatus(req.Status)
	if req.Status == string(models.MissionStatusApproved) {
		mission.ApprovedAt = &now
		mission.SupervisorID = &userID
	} else if req.Status == string(models.MissionStatusCompleted) {
		mission.CompletedAt = &now
	}

	mission, err = s.MissionRepository.Save(mission)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	// Notificar cambio de estado
	s.NotifyWebSocket("mission_status_changed", map[string]interface{}{
		"mission_id":   mission.ID,
		"status":       mission.Status,
		"alchemist_id": mission.AlchemistID,
	})

	response := api.MissionResponse{
		ID:          mission.ID,
		Title:       mission.Title,
		Description: mission.Description,
		Status:      string(mission.Status),
		AlchemistID: mission.AlchemistID,
		RequestedAt: mission.RequestedAt.Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleDeleteMission(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mission, err := s.MissionRepository.FindById(uint(id))
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	if mission == nil {
		http.Error(w, "Mission not found", http.StatusNotFound)
		return
	}

	if err := s.MissionRepository.Delete(mission); err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
