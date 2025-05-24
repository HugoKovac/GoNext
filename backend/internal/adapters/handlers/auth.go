// internal/adapters/handlers/auth_handler.go
package handlers

import (
	"GoNext/base/internal/adapters/dto"
	"GoNext/base/internal/core/domain"
	"GoNext/base/internal/core/ports"
	customvalidator "GoNext/base/pkg/validator"
	"log"
	"os"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService ports.AuthService
	userService ports.UserService
	validate    *validator.Validate
}

func NewAuthHandler(authService ports.AuthService, userService ports.UserService) *AuthHandler {
	v := validator.New()
	customvalidator.RegisterCustomValidators(v)

	return &AuthHandler{
		authService: authService,
		userService: userService,
		validate:    v,
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	// Validate credentials
	if err := h.validate.Struct(dto.UserCredentials{
		Email:    user.Email,
		Password: user.Password,
	}); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	_, err := h.userService.Register(user)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := h.authService.Authenticate(user.Email, user.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// Don't return the password
	user.Password = "********"

	var domain string
	if os.Getenv("DEV") == "true" {
		domain = ".localhost"
	} else {
		domain = ".gonext.com"
	}
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    token,
		Domain:   domain,
		Secure:   true,
		HTTPOnly: true,
		SameSite: "Lax",
		Expires:  time.Now().Add(time.Hour * 24),
		Path:     "/",
	}
	c.Cookie(&cookie)

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var creds dto.UserCredentials
	if err := c.BodyParser(&creds); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	if err := h.validate.Struct(creds); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString("Validation Error")
	}

	token, err := h.authService.Authenticate(creds.Email, creds.Password)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	var domain string
	if os.Getenv("DEV") == "true" {
		domain = ".localhost"
	} else {
		domain = ".gonext.com"
	}
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    token,
		Domain:   domain,
		Secure:   true,
		HTTPOnly: true,
		SameSite: "Lax",
		Expires:  time.Now().Add(time.Hour * 24),
		Path:     "/",
	}
	c.Cookie(&cookie)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
	})
}

func (h *AuthHandler) Status(c *fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "No token provided",
		})
	}

	claims, err := h.authService.ValidateToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	return c.JSON(fiber.Map{
		"user": claims,
	})
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	var domain string
	if os.Getenv("DEV") == "true" {
		domain = ".localhost"
	} else {
		domain = ".gonext.com"
	}
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    "",
		Domain:   domain,
		Secure:   true,
		HTTPOnly: true,
		SameSite: "Lax",
		Expires:  time.Now().Add(-time.Hour), // Past time to expire the cookie
		Path:     "/",
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Logged out successfully",
	})
}
