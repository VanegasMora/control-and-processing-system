package models

import (
	"time"

	"gorm.io/gorm"
)

type MaterialType string

const (
	MaterialTypeMetal     MaterialType = "metal"
	MaterialTypeMineral   MaterialType = "mineral"
	MaterialTypeOrganic   MaterialType = "organic"
	MaterialTypeSynthetic MaterialType = "synthetic"
)

type Material struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Name        string         `gorm:"not null;uniqueIndex" json:"name"`
	Type        MaterialType   `gorm:"not null" json:"type"`
	Description string         `gorm:"type:text" json:"description"`
	Stock       float64        `gorm:"default:0" json:"stock"`
	Unit        string         `gorm:"default:'kg'" json:"unit"`
	Price       float64        `gorm:"default:0" json:"price"`
}
