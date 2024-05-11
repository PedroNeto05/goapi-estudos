package helpers

import (
	"goApi/db/mysql"
	"goApi/utils"

	"gorm.io/gorm"
)

var (
	logger *utils.Logger
	db     *gorm.DB
)

func InitializeHelpers() {
	logger = utils.GetLogger("helpers")
	db, _ = mysql.InitializeMysql()
}
