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

func (s *Server) HandleAlchemists(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetAllAlchemists(w, r)
	case http.MethodPost:
		s.handleCreateAlchemist(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) HandleAlchemistByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetAlchemistByID(w, r)
	case http.MethodPut:
		s.handleUpdateAlchemist(w, r)
	case http.MethodDelete:
		s.handleDeleteAlchemist(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleGetAllAlchemists(w http.ResponseWriter, r *http.Request) {
	alchemists, err := s.AlchemistRepository.FindAll()
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	response := make([]api.AlchemistResponse, len(alchemists))
	for i, a := range alchemists {
		response[i] = api.AlchemistResponse{
			ID:        a.ID,
			Name:      a.Name,
			Email:     a.Email,
			Rank:      string(a.Rank),
			Specialty: string(a.Specialty),
			Role:      string(a.Role),
			Certified: a.Certified,
			CreatedAt: a.CreatedAt.Format(time.RFC3339),
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleGetAlchemistByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	alchemist, err := s.AlchemistRepository.FindById(uint(id))
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	if alchemist == nil {
		http.Error(w, "Alchemist not found", http.StatusNotFound)
		return
	}

	response := api.AlchemistResponse{
		ID:        alchemist.ID,
		Name:      alchemist.Name,
		Email:     alchemist.Email,
		Rank:      string(alchemist.Rank),
		Specialty: string(alchemist.Specialty),
		Role:      string(alchemist.Role),
		Certified: alchemist.Certified,
		CreatedAt: alchemist.CreatedAt.Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleCreateAlchemist(w http.ResponseWriter, r *http.Request) {
	var req api.AlchemistRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Verificar si el email ya existe
	existing, err := s.AlchemistRepository.FindByEmail(req.Email)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	if existing != nil {
		http.Error(w, "Email already registered", http.StatusConflict)
		return
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	role := models.RoleAlchemist
	if req.Role != "" {
		role = models.UserRole(req.Role)
	}

	alchemist := &models.Alchemist{
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashedPassword,
		Rank:      models.AlchemistRank(req.Rank),
		Specialty: models.AlchemistSpecialty(req.Specialty),
		Role:      role,
		Certified: req.Certified,
	}

	alchemist, err = s.AlchemistRepository.Create(alchemist)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	response := api.AlchemistResponse{
		ID:        alchemist.ID,
		Name:      alchemist.Name,
		Email:     alchemist.Email,
		Rank:      string(alchemist.Rank),
		Specialty: string(alchemist.Specialty),
		Role:      string(alchemist.Role),
		Certified: alchemist.Certified,
		CreatedAt: alchemist.CreatedAt.Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleUpdateAlchemist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	alchemist, err := s.AlchemistRepository.FindById(uint(id))
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	if alchemist == nil {
		http.Error(w, "Alchemist not found", http.StatusNotFound)
		return
	}

	var req api.AlchemistRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	alchemist.Name = req.Name
	alchemist.Rank = models.AlchemistRank(req.Rank)
	alchemist.Specialty = models.AlchemistSpecialty(req.Specialty)
	alchemist.Certified = req.Certified

	if req.Password != "" {
		hashedPassword, err := auth.HashPassword(req.Password)
		if err != nil {
			s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
			return
		}
		alchemist.Password = hashedPassword
	}

	alchemist, err = s.AlchemistRepository.Save(alchemist)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	response := api.AlchemistResponse{
		ID:        alchemist.ID,
		Name:      alchemist.Name,
		Email:     alchemist.Email,
		Rank:      string(alchemist.Rank),
		Specialty: string(alchemist.Specialty),
		Role:      string(alchemist.Role),
		Certified: alchemist.Certified,
		CreatedAt: alchemist.CreatedAt.Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleDeleteAlchemist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	alchemist, err := s.AlchemistRepository.FindById(uint(id))
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	if alchemist == nil {
		http.Error(w, "Alchemist not found", http.StatusNotFound)
		return
	}

	if err := s.AlchemistRepository.Delete(alchemist); err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
