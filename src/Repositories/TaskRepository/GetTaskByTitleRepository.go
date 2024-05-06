package taskrepository

import (
	models "goApi/src/Models"

	"gorm.io/gorm"
)

func GetTaskByTitleRepository(title string) *gorm.DB {
	task := db.First(&models.Task{})
	return task
}
