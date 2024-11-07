package models

import (
	"time"

	"github.com/google/uuid"
)

type Coment struct {
	ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Body      string     `gorm:"type:text;not null"`
	CreatedAt *time.Time `gorm:"not null;default:now()"`
	UpdatedAt *time.Time `gorm:"not null;default:now()"`
	FormulaID *uuid.UUID
}

type ComentInput struct {
	Body string `json:"body" validate:"required"`
}
