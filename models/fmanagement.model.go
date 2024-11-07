package models

import (
	"time"

	"github.com/google/uuid"
)

type FormulaManagement struct {
	ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string     `gorm:"type:varchar(100);unique;not null"`
	Status    string     `gorm:"type:varchar(100);not null"`
	Version   string     `gorm:"type:varchar(100);not null"`
	CreatedAt *time.Time `gorm:"not null;default:now()"`
	UpdatedAt *time.Time `gorm:"not null;default:now()"`
	Formulas  []Formula
	ProjectID *uuid.UUID
}

type FormulaManagementInput struct {
	Name    string `json:"name" validate:"required"`
	Status  string `json:"status" validate:"required"`
	Version string `json:"version" validate:"required"`
}
