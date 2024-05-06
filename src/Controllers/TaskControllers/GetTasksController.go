package taskcontrollers

import "github.com/gofiber/fiber/v2"

func GetTasksController(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"success": "INDEX"})
}
