package router

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/zulkarnen-force/fiber-starter/controller"
	"github.com/zulkarnen-force/fiber-starter/repository/postgresql"
	"github.com/zulkarnen-force/fiber-starter/usecase"
)

func Setup(app *fiber.App) {
	userUsecase := usecase.NewUserUsecase(postgresql.NewUserRepository(),"your_jwt_secret", time.Hour*24);
	v1Group := app.Group("/v1")
	controller.NewUserController(userUsecase).Route(v1Group)

	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: []byte("your_jwt_secret"),
	// }))
}
