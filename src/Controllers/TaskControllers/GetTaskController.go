package taskcontrollers

import "github.com/gofiber/fiber/v2"

func GetTaskController(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"success": "SHOW"})
}
