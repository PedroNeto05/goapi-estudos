package userhandler

import (
	"goApi/utils"
)

var (
	logger *utils.Logger
)

func InitializeUserHandler() {
	logger = utils.GetLogger("user_handler")
}
