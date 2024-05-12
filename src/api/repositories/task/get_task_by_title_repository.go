package taskrepository

import (
	"goApi/db/models"

	"github.com/google/uuid"
)

func GetTaskByTitleRepository(title string, userID uuid.UUID) (*models.Task, error) {
	var task models.Task

	err := db.Where("title = ? AND user_id = ?", title, userID).First(&task).Error

	if err != nil {
		return nil, err
	}

	return &task, nil
}
