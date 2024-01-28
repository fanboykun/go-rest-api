package models

import (
	"github.com/jinzhu/gorm"
)

type Album struct {
	gorm.Model
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}