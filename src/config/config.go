package config

import (
	"fmt"
	"goApi/db/mysql"
)

func Init() error {
	_, err := mysql.InitializeMysql()

	if err != nil {
		return fmt.Errorf("erro ao inicializar a base de dados: %v", err)
	}

	return err
}
