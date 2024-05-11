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
		logger.Errorf("Erro ao fazer o parse das informações: %v", err)
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	tasks := taskusecase.GetTasksUseCase(userIdUUID)

	return c.JSON(fiber.Map{"tasks": tasks})
}
