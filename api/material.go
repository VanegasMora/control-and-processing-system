package api

type MaterialResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Stock       float64 `json:"stock"`
	Unit        string  `json:"unit"`
	Price       float64 `json:"price"`
}

type MaterialRequest struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Stock       float64 `json:"stock"`
	Unit        string  `json:"unit"`
	Price       float64 `json:"price"`
}
