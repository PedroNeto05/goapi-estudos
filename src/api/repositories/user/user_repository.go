package userrepository

import (
	"goApi/db/mysql"
	"goApi/utils"

	"gorm.io/gorm"
)

var (
	logger *utils.Logger
	db     *gorm.DB
)

func InitializeUserRepository() {
	logger = utils.GetLogger("user_repository")
	db, _ = mysql.InitializeMysql()
}
