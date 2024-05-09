package userhandler

import (
	"fmt"
	userusecase "goApi/api/usecase/user"
	"goApi/db/models"

	"github.com/gofiber/fiber/v2"
)

func CreateUserHandler(c *fiber.Ctx) error {
	var user *models.User

	err := c.BodyParser(user)

	if err != nil {
		logger.Errorf("Erro ao fazer o parse das informações: %v", err)
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	fmt.Println("sdpkaspidhgaspidh")

	err = userusecase.CreateUserUseCase(user)

	if err != nil {
		logger.Errorf("Erro na criação do usuario: %v", err)
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}
