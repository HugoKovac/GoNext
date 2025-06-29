package handlers

import (
	"GoNext/base/internal/core/ports"
	customvalidator "GoNext/base/pkg/validator"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService ports.UserService
	validate    *validator.Validate
}

func NewUserHandler(service ports.UserService) *UserHandler {
	v := validator.New()
	customvalidator.RegisterCustomValidators(v)

	return &UserHandler{
		userService: service,
		validate:    v,
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

func (h *UserHandler) UpdateCurrentUser(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	var userUpdate struct {
		Email       string `json:"email" validate:"email"`
		OldPassword string `json:"oldPassword" validate:"password"`
		NewPassword string `json:"newPassword" validate:"password"`
	}

	h.validate.Struct(&userUpdate)

	if err := c.BodyParser(&userUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user, err := h.userService.Update(userID, userUpdate.Email, userUpdate.OldPassword, userUpdate.NewPassword)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}
	user.Password = "" // Don't return the password
	return c.JSON(user)
}
