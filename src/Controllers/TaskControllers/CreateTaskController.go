package taskcontrollers

import (
	models "goApi/src/Models"
	taskusecase "goApi/src/UseCase/TaskUseCase"

	"github.com/gofiber/fiber/v2"
)

func CreateTaskController(c *fiber.Ctx) error {
	var myJson models.Task

	err := c.BodyParser(&myJson)

	if err != nil {
		logger.Errorf("Erro ao fazer o parse das informações: %v", err)
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})

	}
	err = taskusecase.CreateTaskUseCase(&myJson)

	if err != nil {
		logger.Errorf("Erro de validação: %v", err)
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}
