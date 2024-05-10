package taskrepository

import (
	"errors"
	"goApi/db/models"

	"gorm.io/gorm"
)

func GetTaskByTitleRepository(title string) (*models.Task, error) {
	var task models.Task
	err := db.Where("title = ?", title).First(&task).Error

	if errors.Is(err, gorm.ErrRecordNotFound) || err != nil {
		return nil, err
	}

	return &task, nil
}
