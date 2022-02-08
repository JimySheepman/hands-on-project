package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	// Adds some metadata fields to the table
	gorm.Model
	// Explicitly specify the type to be uuid
	ID       uuid.UUID `gorm:"type:uuid"`
	Title    string
	SubTitle string
	Text     string
}
