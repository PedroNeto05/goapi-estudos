package main

import (
	"goApi/api/router"
	"goApi/config"
)

func main() {

	err := config.Init()

	if err != nil {
		return
	}

	router.Initialize()
}
