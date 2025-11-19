package auth

import (
	"net/http"
	"strings"

	"github.com/gorilla/context"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		token := parts[1]
		claims, err := ValidateToken(token)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		context.Set(r, "user_id", claims.UserID)
		context.Set(r, "user_email", claims.Email)
		context.Set(r, "user_role", claims.Role)

		next.ServeHTTP(w, r)
	})
}

func SupervisorOnlyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, ok := context.Get(r, "user_role").(string)
		if !ok || role != "supervisor" {
			http.Error(w, "Supervisor access required", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetUserID(r *http.Request) (uint, bool) {
	userID, ok := context.Get(r, "user_id").(uint)
	return userID, ok
}

func GetUserRole(r *http.Request) (string, bool) {
	role, ok := context.Get(r, "user_role").(string)
	return role, ok
}
