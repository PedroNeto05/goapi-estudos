package taskrepository

import (
	"goApi/config"

	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeTaskRepository() {
	logger = config.GetLogger("TaskRepository")
	db, _ = config.InitializeMysql()
}
