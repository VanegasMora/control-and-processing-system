package models

import (
	"time"

	"gorm.io/gorm"
)

type AuditType string

const (
	AuditTypeMaterialUsage AuditType = "material_usage"
	AuditTypeMissionCheck  AuditType = "mission_check"
	AuditTypeTransmutation AuditType = "transmutation"
	AuditTypeSystem        AuditType = "system"
)

type AuditSeverity string

const (
	AuditSeverityLow      AuditSeverity = "low"
	AuditSeverityMedium   AuditSeverity = "medium"
	AuditSeverityHigh     AuditSeverity = "high"
	AuditSeverityCritical AuditSeverity = "critical"
)

type Audit struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Type        AuditType      `gorm:"not null" json:"type"`
	Severity    AuditSeverity  `gorm:"not null;default:'low'" json:"severity"`
	Description string         `gorm:"type:text" json:"description"`
	AlchemistID *uint          `json:"alchemist_id,omitempty"`
	Alchemist   *Alchemist     `gorm:"foreignKey:AlchemistID" json:"alchemist,omitempty"`
	Details     string         `gorm:"type:jsonb" json:"details,omitempty"` // Para almacenar datos adicionales en formato JSON
	Resolved    bool           `gorm:"default:false" json:"resolved"`
	ResolvedAt  *time.Time     `json:"resolved_at,omitempty"`
	ResolvedBy  *uint          `json:"resolved_by,omitempty"`
}
