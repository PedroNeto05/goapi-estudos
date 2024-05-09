package taskhandlers

import (
	taskusecase "goApi/api/usecase/task"
	"goApi/utils"
)

var (
	logger *utils.Logger
)

func InitializeTaskHandler() {
	logger = utils.GetLogger("TaskController")

	taskusecase.InitializeTaskUseCase()
}
