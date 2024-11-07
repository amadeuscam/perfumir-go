package models

import (
	"time"

	"github.com/google/uuid"
)

type Formula struct {
	ID                  *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name                string     `gorm:"type:varchar(255);uniqueIndex;not null"`
	Status              string     `gorm:"type:varchar(100);not null"`
	Source              string     `gorm:"type:varchar(255);not null"`
	Version             string     `gorm:"type:varchar(100);unique;not null"`
	CreatedAt           *time.Time `gorm:"not null;default:now()"`
	UpdatedAt           *time.Time `gorm:"not null;default:now()"`
	FormulaManagementID *uuid.UUID
	Coments             []Coment
	FormulaIngredients  []FormulaIngredient
}

type FormulaInput struct {
	Name    string `json:"name" validate:"required"`
	Status  string `json:"status" validate:"required"`
	Source  string `json:"source" validate:"required"`
	Version string `json:"version" validate:"required"`
}
