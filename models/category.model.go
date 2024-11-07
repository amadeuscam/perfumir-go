package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID           *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name         string     `gorm:"type:varchar(255);uniqueIndex;not null"`
	CreatedAt    *time.Time `gorm:"not null;default:now()"`
	UpdatedAt    *time.Time `gorm:"not null;default:now()"`
	IngredientID *uuid.UUID
}


