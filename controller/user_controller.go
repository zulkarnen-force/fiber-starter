package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zulkarnen-force/fiber-starter/domain"
)

type UserController struct {
	usecase domain.UserUsecase
}


func (h UserController) Register(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.usecase.Register(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(user.ToJson())
}

func (h UserController) Login(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	token, err := h.usecase.Login(user.Email, user.Password);

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"token": token})
}


func NewUserController(usecase domain.UserUsecase) *UserController {
	return &UserController{usecase}
}

func (c UserController) Route(r fiber.Router) {
	r.Post("/register", c.Register);
	r.Post("/login", c.Login)
}
