package taskcontrollers

import "github.com/gofiber/fiber/v2"

func DeleteTaskController(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"success": "DELETE"})
}
