package api

type AuditResponse struct {
	ID          uint               `json:"id"`
	Type        string             `json:"type"`
	Severity    string             `json:"severity"`
	Description string             `json:"description"`
	AlchemistID *uint              `json:"alchemist_id,omitempty"`
	Alchemist   *AlchemistResponse `json:"alchemist,omitempty"`
	Details     string             `json:"details,omitempty"`
	Resolved    bool               `json:"resolved"`
	ResolvedAt  *string            `json:"resolved_at,omitempty"`
	ResolvedBy  *uint              `json:"resolved_by,omitempty"`
	CreatedAt   string             `json:"created_at"`
}

type AuditRequest struct {
	Type        string `json:"type"`
	Severity    string `json:"severity"`
	Description string `json:"description"`
	AlchemistID *uint  `json:"alchemist_id,omitempty"`
	Details     string `json:"details,omitempty"`
}
