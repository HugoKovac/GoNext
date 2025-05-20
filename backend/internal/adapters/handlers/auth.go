// internal/adapters/handlers/auth_handler.go
package handlers

import (
	"GoNext/base/internal/core/domain"
	"GoNext/base/internal/core/ports"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService ports.AuthService
	userService ports.UserService
}

func NewAuthHandler(authService ports.AuthService, userService ports.UserService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		userService: userService,
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate user data
	if user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email, and password are required",
		})
	}

	userCopy := user

	dUser, err := h.userService.Register(&user)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

    creds := domain.UserCredentials{
        Email: userCopy.Email,
        Password: userCopy.Password,
    }

    token, err := h.authService.Authenticate(creds)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// Don't return the password
	dUser.Password = ""

    cookie := fiber.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24),
	}
	c.Cookie(&cookie)

	return c.Status(fiber.StatusCreated).JSON(dUser)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var creds domain.UserCredentials
	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate credentials
	if creds.Email == "" || creds.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email and password are required",
		})
	}

	token, err := h.authService.Authenticate(creds)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	cookie := fiber.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24),
	}
	c.Cookie(&cookie)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
	})
}
