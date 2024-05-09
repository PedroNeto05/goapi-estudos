package taskhandlers

import "github.com/gofiber/fiber/v2"

func DeleteTaskHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"success": "DELETE"})
}
