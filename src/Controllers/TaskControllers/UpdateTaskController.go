package taskcontrollers

import "github.com/gofiber/fiber/v2"

func UpdateTaskController(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"success": "UPDATE"})
}
