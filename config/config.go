package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	_, err := InitializeMysql()

	if err != nil {
		return fmt.Errorf("erro ao inicializar a base de dados: %v", err)
	}

	return err
}

func GetLogger(p string) *Logger {
	logger := setUpLogger(p)
	return logger
}
