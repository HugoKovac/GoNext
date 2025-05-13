package handlers

import (
	"GoNext/base/internal/core/domain"
	"GoNext/base/internal/core/ports"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *ports.UserService
}

func NewUserHandler(service *ports.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	fmt.Println(user)

	return c.JSON("Create User")
}

