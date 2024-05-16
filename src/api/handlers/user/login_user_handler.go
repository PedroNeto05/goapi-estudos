package userhandler

import (
	userusecase "goApi/api/usecase/user"
	"goApi/db/models"

	"github.com/gofiber/fiber/v2"
)

func LoginUserHandler(c *fiber.Ctx) error {

	var loginRequest *models.LoginRequest

	err := c.BodyParser(&loginRequest)

	if err != nil {
		logger.Errorf("Error parsing information: %v", err)
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err,
		})
	}

	login := models.LoginRequest{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}

	token, err := userusecase.LoginUserUseCase(&login)

	if err != nil {
		logger.Errorf("Erro ao fazer o login: %v", err)
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err,
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:  "token.connect",
		Value: *token,
	})

	return c.SendString("logando com sucesso")
}
