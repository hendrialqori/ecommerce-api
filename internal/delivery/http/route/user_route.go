package route

import (
	"internship-mini-project/internal/delivery/http/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoute(app *fiber.App, userHandler *handler.UserHandler) {
	route := app.Group("/api/v1/user")
	route.Post("/register", userHandler.Register)
	route.Post("/login", userHandler.Login)
}
