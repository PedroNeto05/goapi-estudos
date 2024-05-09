package mysql

import (
	"goApi/db/models"
	"goApi/utils"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	logger *utils.Logger
)

func InitializeMysql() (*gorm.DB, error) {
	logger = utils.GetLogger("mysql")

	err := godotenv.Load("../.env")

	if err != nil {
		logger.Errorf("erro ao ler o arquivo .env: %v", err)
		return nil, err
	}

	dbUlr := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(mysql.Open(dbUlr), &gorm.Config{})

	if err != nil {
		logger.Errorf("Erro ao inicializar o Mysql: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&models.Task{})

	if err != nil {
		logger.Errorf("Erro ao fazer as migrações: %v", err)
		return nil, err
	}

	return db, nil
}
