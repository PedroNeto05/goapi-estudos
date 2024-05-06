package taskcontrollers

import (
	"fmt"
	models "goApi/src/Models"

	"github.com/gofiber/fiber/v2"
)

func CreateTaskController(c *fiber.Ctx) error {
	var myJson models.Task

	err := c.BodyParser(&myJson)

	fmt.Printf("%T", myJson)

	// err = taskusecase.CreateTaskUseCase()
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err,
		})
	}

	return c.SendStatus(fiber.StatusOK)
}
