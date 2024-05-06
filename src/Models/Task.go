package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"descripiton"`
}

type TaskResponse struct {
	Id          uint64    `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deteledAt,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"descripiton"`
}
