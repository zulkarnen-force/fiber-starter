package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zulkarnen-force/fiber-starter/domain"
)

type UserController struct {
	usecase domain.UserUsecase
}


func (h *UserController) Register(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.usecase.Register(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

func NewUserController(usecase domain.UserUsecase) *UserController {
	return &UserController{usecase}
}

func (c UserController) Route(r fiber.Router) {
	r.Get("/register", c.Register)
}
