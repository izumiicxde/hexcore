package middleware

import (
	"hexcore/utils"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the token from cookies
		tokenString := c.Cookies("token") // Change this to match your cookie name
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing token"})
		}

		// Verify the token
		_, claims, err := utils.ParseJWT(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
		}

		// Extract user ID from claims
		userID, ok := claims["userId"].(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token data"})
		}

		// Attach user ID to context
		c.Locals("userID", uint(userID))

		// Proceed to next middleware/handler
		return c.Next()
	}
}
