package taskcontrollers

import (
	"goApi/config"
	taskusecase "goApi/src/UseCase/TaskUseCase"
)

var (
	logger *config.Logger
)

func InitializeTaskControler() {
	logger = config.GetLogger("TaskController")

	taskusecase.InitializeTaskUseCase()
}
