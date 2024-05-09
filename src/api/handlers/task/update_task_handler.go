package taskhandlers

import "github.com/gofiber/fiber/v2"

func UpdateTaskHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"success": "UPDATE"})
}
