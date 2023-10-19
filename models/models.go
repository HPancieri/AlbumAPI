package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Album struct {
	gorm.Model
	ID     uuid.UUID `json:"ID" gorm:"uuid;not null;default:uuid_generate_v4"`
	Title  string    `json:"title" gorm:"text;not null;default:null"`
	Artist string    `json:"artist" gorm:"text;not null;default:null"`
	Price  float64   `json:"price" gorm:"float;not null;default null"`
}
