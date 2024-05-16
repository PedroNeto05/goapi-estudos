package taskhandlers

import (
	taskusecase "goApi/api/usecase/task"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func DeleteTaskHandler(c *fiber.Ctx) error {

	taskID := c.Params("taskId")

	userID := c.Locals("userID").(string)

	userIdUUID, err := uuid.Parse(userID)

	if err != nil {
		logger.Errorf("Error deleting the task: %v", err)
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = taskusecase.DeleteTaskUseCase(taskID, userIdUUID)

	if err != nil {
		logger.Errorf("Error deleting the task: %v", err)
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}
