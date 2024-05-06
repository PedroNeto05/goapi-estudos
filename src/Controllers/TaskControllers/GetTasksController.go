package taskcontrollers

import (
	taskusecase "goApi/src/UseCase/TaskUseCase"

	"github.com/gofiber/fiber/v2"
)

func GetTasksController(c *fiber.Ctx) error {

	tasks := taskusecase.GetTasksUseCase()

	return c.JSON(fiber.Map{"takss": tasks})
}
