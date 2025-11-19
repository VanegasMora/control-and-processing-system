package server

import (
	"backend-avanzada/api"
	"backend-avanzada/auth"
	"backend-avanzada/models"
	"encoding/json"
	"net/http"
	"time"
)

func (s *Server) HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req api.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Formato de solicitud inválido"})
		return
	}

	alchemist, err := s.AlchemistRepository.FindByEmail(req.Email)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	if alchemist == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Usuario no encontrado"})
		return
	}

	if !auth.CheckPasswordHash(req.Password, alchemist.Password) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Contraseña incorrecta"})
		return
	}

	token, err := auth.GenerateToken(alchemist.ID, alchemist.Email, string(alchemist.Role))
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	response := api.AuthResponse{
		Token:     token,
		UserID:    alchemist.ID,
		Email:     alchemist.Email,
		Role:      string(alchemist.Role),
		Name:      alchemist.Name,
		Rank:      string(alchemist.Rank),
		Specialty: string(alchemist.Specialty),
	}

	// Los headers CORS ya están establecidos por el middleware
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (s *Server) HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req api.RegisterRequest
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

	alchemist := &models.Alchemist{
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashedPassword,
		Rank:      models.AlchemistRank(req.Rank),
		Specialty: models.AlchemistSpecialty(req.Specialty),
		Role:      models.RoleAlchemist,
		Certified: false,
	}

	alchemist, err = s.AlchemistRepository.Create(alchemist)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	token, err := auth.GenerateToken(alchemist.ID, alchemist.Email, string(alchemist.Role))
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	response := api.AuthResponse{
		Token:     token,
		UserID:    alchemist.ID,
		Email:     alchemist.Email,
		Role:      string(alchemist.Role),
		Name:      alchemist.Name,
		Rank:      string(alchemist.Rank),
		Specialty: string(alchemist.Specialty),
	}

	// Los headers CORS ya están establecidos por el middleware
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (s *Server) HandleGetProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := auth.GetUserID(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	alchemist, err := s.AlchemistRepository.FindById(userID)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	if alchemist == nil {
		http.Error(w, "User not found", http.StatusNotFound)
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
