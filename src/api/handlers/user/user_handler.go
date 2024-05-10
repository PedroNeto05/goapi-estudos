package userhandler

import (
	userusecase "goApi/api/usecase/user"
	"goApi/utils"
)

var (
	logger *utils.Logger
)

func InitializeUserHandler() {
	logger = utils.GetLogger("user_handler")
	userusecase.InitializeUserUseCase()
}
