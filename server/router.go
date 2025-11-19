package server

import (
	"backend-avanzada/auth"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) router() http.Handler {
	router := mux.NewRouter()

	// Logger después de CORS
	router.Use(s.logger.RequestLogger)

	// Rutas públicas (OPTIONS manejado por middleware CORS)
	router.HandleFunc("/api/auth/login", s.HandleLogin).Methods(http.MethodPost)
	router.HandleFunc("/api/auth/register", s.HandleRegister).Methods(http.MethodPost)

	// Handler explícito para OPTIONS en todas las rutas /api/*
	router.PathPrefix("/api/").Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Headers CORS ya establecidos por el middleware
		w.WriteHeader(http.StatusOK)
	})

	// Rutas protegidas
	api := router.PathPrefix("/api").Subrouter()
	api.Use(auth.AuthMiddleware)

	// Perfil
	api.HandleFunc("/auth/profile", s.HandleGetProfile).Methods(http.MethodGet)

	// Alquimistas
	api.HandleFunc("/alchemists", s.HandleAlchemists).Methods(http.MethodGet, http.MethodPost)
	api.HandleFunc("/alchemists/{id}", s.HandleAlchemistByID).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)

	// Misiones
	api.HandleFunc("/missions", s.HandleMissions).Methods(http.MethodGet, http.MethodPost)
	api.HandleFunc("/missions/{id}", s.HandleMissionByID).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)
	api.HandleFunc("/missions/{id}/status", s.HandleUpdateMissionStatus).Methods(http.MethodPut)

	// Materiales
	api.HandleFunc("/materials", s.HandleMaterials).Methods(http.MethodGet, http.MethodPost)
	api.HandleFunc("/materials/{id}", s.HandleMaterialByID).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)

	// Transmutaciones
	api.HandleFunc("/transmutations", s.HandleTransmutations).Methods(http.MethodGet, http.MethodPost)
	api.HandleFunc("/transmutations/{id}", s.HandleTransmutationByID).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)
	api.HandleFunc("/transmutations/{id}/status", s.HandleUpdateTransmutationStatus).Methods(http.MethodPut)

	// Auditorías
	api.HandleFunc("/audits", s.HandleAudits).Methods(http.MethodGet, http.MethodPost)
	api.HandleFunc("/audits/{id}", s.HandleAuditByID).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)
	api.HandleFunc("/audits/{id}/resolve", s.HandleResolveAudit).Methods(http.MethodPut)

	// WebSocket (sin middleware de auth, se valida en el handler)
	router.HandleFunc("/api/ws", s.HandleWebSocket).Methods(http.MethodGet)

	// Aplicar CORS como wrapper externo - maneja OPTIONS antes del router
	return corsMiddleware(router)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		// Si hay un Origin, usarlo; si no, permitir cualquier origen
		if origin == "" {
			origin = "*"
			// Si usamos *, no podemos usar credentials
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			// Origen específico - podemos usar credentials
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Max-Age", "3600")

		// Manejar preflight OPTIONS - responder antes de pasar al router
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
