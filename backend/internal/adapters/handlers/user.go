package handlers

import (
	"GoNext/base/internal/core/ports"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService ports.UserService
}

func NewUserHandler(service ports.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (h *UserHandler) GetById(c *fiber.Ctx) error {
	idParam := struct {
		ID string `json:"id"`
	}{}

	if err := c.BodyParser(&idParam); err != nil { // ? where is validator
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	
	user, err := h.userService.GetById(idParam.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}
	return c.JSON(user)
}

func (h *UserHandler) GetByEmail(c *fiber.Ctx) error {
	emailParam := struct {
		Email string `json:"email"`
	}{}

	if err := c.BodyParser(&emailParam); err != nil { // ? where is validator
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user, err := h.userService.GetByEmail(emailParam.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}
	user.Password = ""
	return c.JSON(user)
}

func (h *UserHandler) GetCurrentUser(c *fiber.Ctx) error {
    userID := c.Locals("userID").(string)

    user, err := h.userService.GetById(userID)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "User not found",
        })
    }
	user.Password = ""

    return c.JSON(user)
}
