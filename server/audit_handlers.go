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

func (s *Server) HandleAudits(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetAllAudits(w, r)
	case http.MethodPost:
		s.handleCreateAudit(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) HandleAuditByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetAuditByID(w, r)
	case http.MethodPut:
		s.handleUpdateAudit(w, r)
	case http.MethodDelete:
		s.handleDeleteAudit(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleGetAllAudits(w http.ResponseWriter, r *http.Request) {
	audits, err := s.AuditRepository.FindAll()
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	response := make([]api.AuditResponse, len(audits))
	for i, a := range audits {
		response[i] = s.auditToResponse(a)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleGetAuditByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	audit, err := s.AuditRepository.FindById(uint(id))
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	if audit == nil {
		http.Error(w, "Audit not found", http.StatusNotFound)
		return
	}

	response := s.auditToResponse(audit)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleCreateAudit(w http.ResponseWriter, r *http.Request) {
	var req api.AuditRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	audit := &models.Audit{
		Type:        models.AuditType(req.Type),
		Severity:    models.AuditSeverity(req.Severity),
		Description: req.Description,
		AlchemistID: req.AlchemistID,
		Details:     req.Details,
		Resolved:    false,
	}

	audit, err := s.AuditRepository.Create(audit)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	response := s.auditToResponse(audit)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (s *Server) HandleResolveAudit(w http.ResponseWriter, r *http.Request) {
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

	audit, err := s.AuditRepository.FindById(uint(id))
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	if audit == nil {
		http.Error(w, "Audit not found", http.StatusNotFound)
		return
	}

	userID, _ := auth.GetUserID(r)
	now := time.Now()

	audit.Resolved = true
	audit.ResolvedAt = &now
	audit.ResolvedBy = &userID

	audit, err = s.AuditRepository.Save(audit)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	response := s.auditToResponse(audit)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleUpdateAudit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	audit, err := s.AuditRepository.FindById(uint(id))
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	if audit == nil {
		http.Error(w, "Audit not found", http.StatusNotFound)
		return
	}

	var req api.AuditRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	audit.Type = models.AuditType(req.Type)
	audit.Severity = models.AuditSeverity(req.Severity)
	audit.Description = req.Description
	audit.Details = req.Details

	audit, err = s.AuditRepository.Save(audit)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	response := s.auditToResponse(audit)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleDeleteAudit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	audit, err := s.AuditRepository.FindById(uint(id))
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	if audit == nil {
		http.Error(w, "Audit not found", http.StatusNotFound)
		return
	}

	if err := s.AuditRepository.Delete(audit); err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) auditToResponse(a *models.Audit) api.AuditResponse {
	var alchemist *api.AlchemistResponse
	if a.Alchemist != nil && a.Alchemist.ID != 0 {
		alchemist = &api.AlchemistResponse{
			ID:        a.Alchemist.ID,
			Name:      a.Alchemist.Name,
			Email:     a.Alchemist.Email,
			Rank:      string(a.Alchemist.Rank),
			Specialty: string(a.Alchemist.Specialty),
			Role:      string(a.Alchemist.Role),
			Certified: a.Alchemist.Certified,
		}
	}

	response := api.AuditResponse{
		ID:          a.ID,
		Type:        string(a.Type),
		Severity:    string(a.Severity),
		Description: a.Description,
		AlchemistID: a.AlchemistID,
		Alchemist:   alchemist,
		Details:     a.Details,
		Resolved:    a.Resolved,
		ResolvedBy:  a.ResolvedBy,
		CreatedAt:   a.CreatedAt.Format(time.RFC3339),
	}

	if a.ResolvedAt != nil {
		resolvedAt := a.ResolvedAt.Format(time.RFC3339)
		response.ResolvedAt = &resolvedAt
	}

	return response
}
