package taskrepository

import (
	"goApi/db/models"
)

func GetTaskByTitleRepository(title string) (*models.Task, error) {
	var task models.Task

	err := db.Where("title = ?", title).First(&task).Error

	if err != nil {
		return nil, err
	}

	return &task, nil
}
