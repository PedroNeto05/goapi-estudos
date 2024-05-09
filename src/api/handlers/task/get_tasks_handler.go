package taskhandlers

import (
	taskusecase "goApi/api/usecase/task"

	"github.com/gofiber/fiber/v2"
)

func GetTasksHandler(c *fiber.Ctx) error {

	tasks := taskusecase.GetTasksUseCase()

	return c.JSON(fiber.Map{"takss": tasks})
}
