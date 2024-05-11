package taskrepository

import (
	"goApi/db/models"

	"github.com/google/uuid"
)

func GetTasksRepository(userID uuid.UUID) ([]models.Task, error) {
	tasks := []models.Task{}
	err := db.Where("user_id = ?", userID).Find(&tasks).Error

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
