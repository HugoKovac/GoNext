// internal/middleware/jwt_middleware.go
package middleware

import (
	"strings"

	"GoNext/base/internal/core/ports"

	"github.com/gofiber/fiber/v2"
)

func JWTAuthentication(authService ports.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("token")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing authorization header",
			})
		}

		parts := strings.Split(token, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid authorization format",
			})
		}

		userID, err := authService.ValidateToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		// Set user ID in context for use in protected routes
		c.Locals("userID", userID)

		return c.Next()
	}
}
