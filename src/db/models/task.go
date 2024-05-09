package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
