package models

import (
	"time"

	"gorm.io/gorm"
)

type TransmutationStatus string

const (
	TransmutationStatusPending   TransmutationStatus = "pending"
	TransmutationStatusApproved  TransmutationStatus = "approved"
	TransmutationStatusRejected  TransmutationStatus = "rejected"
	TransmutationStatusCompleted TransmutationStatus = "completed"
	TransmutationStatusFailed    TransmutationStatus = "failed"
)

type Transmutation struct {
	ID              uint                    `gorm:"primaryKey" json:"id"`
	CreatedAt       time.Time               `json:"created_at"`
	UpdatedAt       time.Time               `json:"updated_at"`
	DeletedAt       gorm.DeletedAt          `gorm:"index" json:"deleted_at,omitempty"`
	AlchemistID     uint                    `gorm:"not null" json:"alchemist_id"`
	Alchemist       Alchemist               `gorm:"foreignKey:AlchemistID" json:"alchemist,omitempty"`
	Status          TransmutationStatus     `gorm:"not null;default:'pending'" json:"status"`
	InputMaterials  []TransmutationMaterial `gorm:"foreignKey:TransmutationID;constraint:OnDelete:CASCADE" json:"input_materials,omitempty"`
	OutputMaterials []TransmutationMaterial `gorm:"foreignKey:TransmutationID;constraint:OnDelete:CASCADE" json:"output_materials,omitempty"`
	Description     string                  `gorm:"type:text" json:"description"`
	Cost            float64                 `gorm:"default:0" json:"cost"`
	Result          string                  `gorm:"type:text" json:"result,omitempty"`
	SupervisorID    *uint                   `json:"supervisor_id,omitempty"`
	ApprovedAt      *time.Time              `json:"approved_at,omitempty"`
	CompletedAt     *time.Time              `json:"completed_at,omitempty"`
}

type TransmutationMaterial struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	TransmutationID uint      `gorm:"not null" json:"transmutation_id"`
	MaterialID      uint      `gorm:"not null" json:"material_id"`
	Material        Material  `gorm:"foreignKey:MaterialID" json:"material,omitempty"`
	Quantity        float64   `gorm:"not null" json:"quantity"`
	IsInput         bool      `gorm:"not null" json:"is_input"` // true para materiales de entrada, false para salida
}
