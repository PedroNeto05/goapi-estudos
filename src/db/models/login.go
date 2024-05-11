package models

import (
	"gorm.io/gorm"
)

type LoginRequest struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}
