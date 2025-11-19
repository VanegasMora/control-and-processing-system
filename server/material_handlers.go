package server

import (
	"backend-avanzada/api"
	"backend-avanzada/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) HandleMaterials(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetAllMaterials(w, r)
	case http.MethodPost:
		s.handleCreateMaterial(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) HandleMaterialByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleGetMaterialByID(w, r)
	case http.MethodPut:
		s.handleUpdateMaterial(w, r)
	case http.MethodDelete:
		s.handleDeleteMaterial(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleGetAllMaterials(w http.ResponseWriter, r *http.Request) {
	materials, err := s.MaterialRepository.FindAll()
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	response := make([]api.MaterialResponse, len(materials))
	for i, m := range materials {
		response[i] = api.MaterialResponse{
			ID:          m.ID,
			Name:        m.Name,
			Type:        string(m.Type),
			Description: m.Description,
			Stock:       m.Stock,
			Unit:        m.Unit,
			Price:       m.Price,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleGetMaterialByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	material, err := s.MaterialRepository.FindById(uint(id))
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	if material == nil {
		http.Error(w, "Material not found", http.StatusNotFound)
		return
	}

	response := api.MaterialResponse{
		ID:          material.ID,
		Name:        material.Name,
		Type:        string(material.Type),
		Description: material.Description,
		Stock:       material.Stock,
		Unit:        material.Unit,
		Price:       material.Price,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleCreateMaterial(w http.ResponseWriter, r *http.Request) {
	var req api.MaterialRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	material := &models.Material{
		Name:        req.Name,
		Type:        models.MaterialType(req.Type),
		Description: req.Description,
		Stock:       req.Stock,
		Unit:        req.Unit,
		Price:       req.Price,
	}

	material, err := s.MaterialRepository.Create(material)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	response := api.MaterialResponse{
		ID:          material.ID,
		Name:        material.Name,
		Type:        string(material.Type),
		Description: material.Description,
		Stock:       material.Stock,
		Unit:        material.Unit,
		Price:       material.Price,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleUpdateMaterial(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	material, err := s.MaterialRepository.FindById(uint(id))
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	if material == nil {
		http.Error(w, "Material not found", http.StatusNotFound)
		return
	}

	var req api.MaterialRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	material.Name = req.Name
	material.Type = models.MaterialType(req.Type)
	material.Description = req.Description
	material.Stock = req.Stock
	material.Unit = req.Unit
	material.Price = req.Price

	material, err = s.MaterialRepository.Save(material)
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	response := api.MaterialResponse{
		ID:          material.ID,
		Name:        material.Name,
		Type:        string(material.Type),
		Description: material.Description,
		Stock:       material.Stock,
		Unit:        material.Unit,
		Price:       material.Price,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleDeleteMaterial(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	material, err := s.MaterialRepository.FindById(uint(id))
	if err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}
	if material == nil {
		http.Error(w, "Material not found", http.StatusNotFound)
		return
	}

	if err := s.MaterialRepository.Delete(material); err != nil {
		s.HandleError(w, http.StatusInternalServerError, r.URL.Path, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
