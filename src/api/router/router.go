package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Initialize() {
	// instnaciando a aplicação
	app := fiber.New()

	// inicializando as rotas
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
	}))

	initializeRoutes(app)

	// app rodando na porta 3000
	app.Listen(":3001")
}
