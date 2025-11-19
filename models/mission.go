package models

import (
	"time"

	"gorm.io/gorm"
)

type MissionStatus string

const (
	MissionStatusPending    MissionStatus = "pending"
	MissionStatusApproved   MissionStatus = "approved"
	MissionStatusInProgress MissionStatus = "in_progress"
	MissionStatusCompleted  MissionStatus = "completed"
	MissionStatusCancelled  MissionStatus = "cancelled"
)

type Mission struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Title        string         `gorm:"not null" json:"title"`
	Description  string         `gorm:"type:text" json:"description"`
	Status       MissionStatus  `gorm:"not null;default:'pending'" json:"status"`
	AlchemistID  uint           `gorm:"not null" json:"alchemist_id"`
	Alchemist    Alchemist      `gorm:"foreignKey:AlchemistID" json:"alchemist,omitempty"`
	RequestedAt  time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"requested_at"`
	ApprovedAt   *time.Time     `json:"approved_at,omitempty"`
	CompletedAt  *time.Time     `json:"completed_at,omitempty"`
	SupervisorID *uint          `json:"supervisor_id,omitempty"`
}
