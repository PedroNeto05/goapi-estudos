package taskrepository

import (
	"goApi/db/models"

	"github.com/google/uuid"
)

func GetTasksRepository(userID uuid.UUID) []models.Task {
	tasks := []models.Task{}
	db.Where("user_id = ?", userID).Find(&tasks)

	return tasks
}
