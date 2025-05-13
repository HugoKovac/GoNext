// internal/middleware/jwt_middleware.go
package middleware

import (
    "strings"

    "github.com/gofiber/fiber/v2"
    "GoNext/base/internal/core/ports"
)

func JWTAuthentication(authService ports.AuthService) fiber.Handler {
    return func(c *fiber.Ctx) error {
        authHeader := c.Get("Authorization")
        if authHeader == "" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Missing authorization header",
            })
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Invalid authorization format",
            })
        }

        userID, err := authService.ValidateToken(parts[1])
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
