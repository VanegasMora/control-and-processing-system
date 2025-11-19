package models

import (
	"time"

	"gorm.io/gorm"
)

type AlchemistRank string

const (
	RankApprentice AlchemistRank = "apprentice"
	RankState      AlchemistRank = "state"
	RankNational   AlchemistRank = "national"
)

type AlchemistSpecialty string

const (
	SpecialtyCombat     AlchemistSpecialty = "combat"
	SpecialtyResearch   AlchemistSpecialty = "research"
	SpecialtyMedical    AlchemistSpecialty = "medical"
	SpecialtyIndustrial AlchemistSpecialty = "industrial"
)

type UserRole string

const (
	RoleAlchemist  UserRole = "alchemist"
	RoleSupervisor UserRole = "supervisor"
)

type Alchemist struct {
	ID             uint               `gorm:"primaryKey" json:"id"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
	DeletedAt      gorm.DeletedAt     `gorm:"index" json:"deleted_at,omitempty"`
	Name           string             `gorm:"not null" json:"name"`
	Email          string             `gorm:"uniqueIndex;not null" json:"email"`
	Password       string             `gorm:"not null" json:"-"` // No exponer en JSON
	Rank           AlchemistRank      `gorm:"not null" json:"rank"`
	Specialty      AlchemistSpecialty `gorm:"not null" json:"specialty"`
	Role           UserRole           `gorm:"not null;default:'alchemist'" json:"role"`
	Certified      bool               `gorm:"default:false" json:"certified"`
	Missions       []Mission          `gorm:"foreignKey:AlchemistID" json:"missions,omitempty"`
	Transmutations []Transmutation    `gorm:"foreignKey:AlchemistID" json:"transmutations,omitempty"`
}
