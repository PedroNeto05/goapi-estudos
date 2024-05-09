package taskusecase

import (
	taskrepository "goApi/api/repositories/TaskRepository"
	"goApi/db/mysql"
	"goApi/utils"

	"gorm.io/gorm"
)

var (
	logger *utils.Logger
	db     *gorm.DB
)

func InitializeTaskUseCase() {
	logger = utils.GetLogger("UseCase")
	db, _ = mysql.InitializeMysql()
	taskrepository.InitializeTaskRepository()
}
