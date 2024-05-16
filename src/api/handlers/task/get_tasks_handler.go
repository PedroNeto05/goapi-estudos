package taskhandlers

import (
	taskusecase "goApi/api/usecase/task"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetTasksHandler(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	userIdUUID, err := uuid.Parse(userID)

	if err != nil {
		logger.Errorf("Error parsing information: %v", err)
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	tasks, err := taskusecase.GetTasksUseCase(userIdUUID)

	if err != nil {
		logger.Errorf("Error fetching tasks: %v", err)
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"tasks": tasks})
}
