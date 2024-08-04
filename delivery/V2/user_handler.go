// delivery/v2/user_handler.go
package v2

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/zulkarnen-force/fiber-starter/domain"
)

type UserHandler struct {
	usecase domain.UserUsecase
}

func NewUserHandler(usecase domain.UserUsecase) *UserHandler {
	return &UserHandler{usecase}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var input struct {
		Name            string `json:"name"`
		Email           string `json:"email"`
		Password        string `json:"password"`
		EmailVerification bool   `json:"email_verification"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user := domain.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	if err := h.usecase.Register(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Assuming EmailVerification is a new feature in v2
	if input.EmailVerification {
		// Handle email verification logic
		// For example, send an email verification link
		// to the user's email address or phone number
		// or store the verification token in the database
		// and associate it with the user
		fmt.Printf("Email verification is not implemented in v2")
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
