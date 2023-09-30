package models

import (
	"gorm.io/gorm"
)

type Album struct {
	gorm.Model
	ID     string  `json:"ID" gorm:"text;not null;default:null"`
	Title  string  `json:"title" gorm:"text;not null;default:null"`
	Artist string  `json:"artist" gorm:"text;not null;default:null"`
	Price  float64 `json:"price" gorm:"float;not null;default null"`
}
