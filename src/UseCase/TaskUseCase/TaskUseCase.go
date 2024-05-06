package taskusecase

import (
	"goApi/config"
	taskrepository "goApi/src/Repositories/TaskRepository"

	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeTaskUseCase() {
	logger = config.GetLogger("UseCase")
	db, _ = config.InitializeMysql()
	taskrepository.InitializeTaskRepository()
}
