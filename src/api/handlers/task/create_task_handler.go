package taskhandlers

import (
	taskusecase "goApi/api/usecase/task"
	"goApi/db/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateTaskHandler(c *fiber.Ctx) error {
	var taskRequest *models.Task

	err := c.BodyParser(&taskRequest)

	if err != nil {
		logger.Errorf("Error parsing information: %v", err)
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userID := c.Locals("userID").(string)

	userIdUUID, err := uuid.Parse(userID)

	if err != nil {
		logger.Errorf("Error parsing information: %v", err)
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	task := models.Task{
		Title:       taskRequest.Title,
		Description: taskRequest.Description,
		UserID:      userIdUUID,
	}

	err = taskusecase.CreateTaskUseCase(&task, userIdUUID)

	if err != nil {
		logger.Errorf("Error creating the task: %v", err)
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}
