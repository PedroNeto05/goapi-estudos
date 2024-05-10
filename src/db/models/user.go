package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email" gorm:"primaryKey"`
	Password string `json:"password"`
}
