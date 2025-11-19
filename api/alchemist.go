package api

type AlchemistResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Rank      string `json:"rank"`
	Specialty string `json:"specialty"`
	Role      string `json:"role"`
	Certified bool   `json:"certified"`
	CreatedAt string `json:"created_at"`
}

type AlchemistRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	Rank      string `json:"rank"`
	Specialty string `json:"specialty"`
	Role      string `json:"role,omitempty"`
	Certified bool   `json:"certified,omitempty"`
}
