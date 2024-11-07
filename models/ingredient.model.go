package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Ingredient struct {
	ID           *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name         string     `gorm:"not null"`
	CasNumber    string     `gorm:"not null"`
	PiramidLevel string     `gorm:"not null"`
	Description  string     `gorm:"type:text;not null"`
	Type         string     `gorm:"type:varchar(255);not null"`
	Ifra         string     `gorm:"type:varchar(255);not null"`
	Inpact       int64
	Life         int64
	Dilutions    pq.Int64Array `gorm:"type:integer[]"`
	CreatedAt    *time.Time    `gorm:"not null;default:now()"`
	UpdatedAt    *time.Time    `gorm:"not null;default:now()"`
	Categorys    []Category
}

type IngredientInput struct {
	Name         string        `json:"name" validate:"required"`
	CasNumber    string        `json:"casNumber" validate:"required"`
	PiramidLevel string        `json:"piramidLevel" validate:"required"`
	Description  string        `json:"description" validate:"required"`
	Type         string        `json:"type" validate:"required"`
	Ifra         string        `json:"ifra" validate:"required"`
	Inpact       int64          `json:"inpact" validate:"required"`
	Life         int64          `json:"life" validate:"required"`
	Dilutions    pq.Int64Array `json:"dilutions" validate:"required"`
}
