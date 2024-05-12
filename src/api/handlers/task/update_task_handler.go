package taskhandlers

import (
	taskusecase "goApi/api/usecase/task"
	"goApi/db/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UpdateTaskHandler(c *fiber.Ctx) error {

	var task *models.Task

	err := c.BodyParser(&task)

	if err != nil {
		logger.Errorf("Erro ao fazer o parse das informações: %v", err)
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userID := c.Locals("userID").(string)

	userIdUUID, err := uuid.Parse(userID)

	taskID := c.Params("taskId")

	if err != nil {
		logger.Errorf("Erro ao fazer o parse das informações: %v", err)
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = taskusecase.UpdateTaskUseCase(task, taskID, userIdUUID)

	if err != nil {
		logger.Errorf("Erro ao fazer o parse das informações: %v", err)
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}
