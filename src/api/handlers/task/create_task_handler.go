package taskhandlers

import (
	taskusecase "goApi/api/usecase/task"
	"goApi/db/models"

	"github.com/gofiber/fiber/v2"
)

func CreateTaskHandler(c *fiber.Ctx) error {
	var task *models.Task

	err := c.BodyParser(&task)

	if err != nil {
		logger.Errorf("Erro ao fazer o parse das informações: %v", err)
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})

	}
	err = taskusecase.CreateTaskUseCase(task)

	if err != nil {
		logger.Errorf("Erro na criação da tarefa: %v", err)
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}
