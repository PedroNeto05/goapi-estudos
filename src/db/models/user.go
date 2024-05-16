package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" gorm:"primaryKey"`
	Password  string `json:"password"`
	Task      []Task `gorm:"foreignKey:UserID"`
}
