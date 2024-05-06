package Router

import "github.com/gofiber/fiber/v2"

func Initialize() {
	// instnaciando a aplicação
	app := fiber.New()

	// inicializando as rotas

	initializeRoutes(app)

	// app rodando na porta 3000
	app.Listen(":3000")
}
