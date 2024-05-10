package taskusecase

import (
	taskrepository "goApi/api/repositories/task"
	"goApi/utils"
)

var (
	logger *utils.Logger
)

func InitializeTaskUseCase() {
	logger = utils.GetLogger("task_usecase")
	taskrepository.InitializeTaskRepository()
}
