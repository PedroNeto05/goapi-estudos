package taskrepository

import (
	models "goApi/src/Models"
)

func GetTaskByTitleRepository(title string) (*models.Task, error) {
	var task models.Task

	err := db.Where("title = ?", title).First(&task).Error

	if err != nil {
		return &task, err
	}
	return &task, nil
}
