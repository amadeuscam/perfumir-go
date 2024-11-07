package models

import (
	"time"

	"github.com/google/uuid"
)

type FormulaIngredient struct {
	ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string     `gorm:"type:varchar(255);not null"`
	Amount    float64    `sql:"type:decimal(8,3);"`
	Dilution  int8
	Alcohol   int8
	CreatedAt *time.Time `gorm:"not null;default:now()"`
	UpdatedAt *time.Time `gorm:"not null;default:now()"`
	FormulaID *uuid.UUID
}

type FormulaIngredientInput struct {
	Name     string `json:"name" validate:"required"`
	Amount   float64 `json:"amount" validate:"required"`
	Dilution int8   `json:"dilution" validate:"required"`
	Alcohol  int8   `json:"alcohol" validate:"required"`
}
