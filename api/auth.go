package api

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Rank      string `json:"rank"`
	Specialty string `json:"specialty"`
}

type AuthResponse struct {
	Token     string `json:"token"`
	UserID    uint   `json:"user_id"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Name      string `json:"name"`
	Rank      string `json:"rank"`
	Specialty string `json:"specialty"`
}
