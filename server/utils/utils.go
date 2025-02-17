package utils

import "github.com/gofiber/fiber/v2"

func WriteJSON(c *fiber.Ctx, status int, data interface{}) error {
	return c.Status(status).JSON(data)
}
func WriteError(c *fiber.Ctx, status int, err error) error {
	return c.Status(status).JSON(map[string]any{"message": err.Error()})
}
