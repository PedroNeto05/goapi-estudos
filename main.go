package main

import (
	"goApi/config"
	"goApi/src/Router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("MAIN")

	err := config.Init()

	if err != nil {
		logger.Errorf("Erro na inicialização da configuração: %v", err)
		return
	}

	Router.Initialize()
}
