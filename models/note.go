package models

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}
