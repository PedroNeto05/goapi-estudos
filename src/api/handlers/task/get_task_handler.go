package taskhandlers

import "github.com/gofiber/fiber/v2"

func GetTaskHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"success": "SHOW"})
}
