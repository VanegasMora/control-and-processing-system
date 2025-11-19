package api

type TransmutationMaterialRequest struct {
	MaterialID uint    `json:"material_id"`
	Quantity   float64 `json:"quantity"`
}

type TransmutationRequest struct {
	Description     string                         `json:"description"`
	InputMaterials  []TransmutationMaterialRequest `json:"input_materials"`
	OutputMaterials []TransmutationMaterialRequest `json:"output_materials,omitempty"`
}

type TransmutationMaterialResponse struct {
	ID       uint              `json:"id"`
	Material *MaterialResponse `json:"material"`
	Quantity float64           `json:"quantity"`
	IsInput  bool              `json:"is_input"`
}

type TransmutationResponse struct {
	ID              uint                            `json:"id"`
	AlchemistID     uint                            `json:"alchemist_id"`
	Alchemist       *AlchemistResponse              `json:"alchemist,omitempty"`
	Status          string                          `json:"status"`
	InputMaterials  []TransmutationMaterialResponse `json:"input_materials,omitempty"`
	OutputMaterials []TransmutationMaterialResponse `json:"output_materials,omitempty"`
	Description     string                          `json:"description"`
	Cost            float64                         `json:"cost"`
	Result          string                          `json:"result,omitempty"`
	SupervisorID    *uint                           `json:"supervisor_id,omitempty"`
	ApprovedAt      *string                         `json:"approved_at,omitempty"`
	CompletedAt     *string                         `json:"completed_at,omitempty"`
	CreatedAt       string                          `json:"created_at"`
}

type TransmutationStatusUpdate struct {
	Status string `json:"status"`
	Result string `json:"result,omitempty"`
}
