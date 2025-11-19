package api

type MissionResponse struct {
	ID           uint               `json:"id"`
	Title        string             `json:"title"`
	Description  string             `json:"description"`
	Status       string             `json:"status"`
	AlchemistID  uint               `json:"alchemist_id"`
	Alchemist    *AlchemistResponse `json:"alchemist,omitempty"`
	RequestedAt  string             `json:"requested_at"`
	ApprovedAt   *string            `json:"approved_at,omitempty"`
	CompletedAt  *string            `json:"completed_at,omitempty"`
	SupervisorID *uint              `json:"supervisor_id,omitempty"`
}

type MissionRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	AlchemistID uint   `json:"alchemist_id,omitempty"`
}

type MissionStatusUpdate struct {
	Status string `json:"status"`
}
