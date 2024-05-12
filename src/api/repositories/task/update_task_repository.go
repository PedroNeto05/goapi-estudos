package taskrepository

import (
	"goApi/db/models"

	"github.com/google/uuid"
)

func UpdateTaskRepository(task *models.Task, taskID string, userID uuid.UUID) error {

	taskParams := models.Task{
		Title:       task.Title,
		Description: task.Description,
	}

	err := db.Model(&models.Task{}).Where("id = ? AND user_id = ?", taskID, userID).Updates(taskParams).Error

	if err != nil {
		return err
	}

	return nil
}
