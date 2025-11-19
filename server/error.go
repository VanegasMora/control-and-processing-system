package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) HandleError(w http.ResponseWriter, status int, path string, err error) {
	s.logger.Error(status, path, err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}
