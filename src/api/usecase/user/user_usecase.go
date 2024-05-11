package userusecase

import (
	"goApi/api/helpers"
	userrepository "goApi/api/repositories/user"
	"goApi/utils"
)

var (
	logger *utils.Logger
)

func InitializeUserUseCase() {
	logger = utils.GetLogger("user_usecase")
	userrepository.InitializeUserRepository()
	helpers.InitializeHelpers()
}
