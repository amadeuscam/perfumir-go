package models

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID                 *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name               string     `gorm:"type:varchar(100);uniqueIndex;not null"`
	CreatedAt          *time.Time `gorm:"not null;default:now()"`
	UpdatedAt          *time.Time `gorm:"not null;default:now()"`
	FormulasManagement []FormulaManagement
}

type ProjectInput struct {
	Name string `json:"name" validate:"required"`
}
