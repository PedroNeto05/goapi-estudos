package userusecase

import (
	userrepository "goApi/api/repositories/user"
	"goApi/utils"
)

var (
	logger *utils.Logger
)

func InitializeUserUseCase() {
	logger = utils.GetLogger("user_usecase")
	userrepository.InitializeUserRepository()
}
