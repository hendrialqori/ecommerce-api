package route

import (
	"internship-mini-project/internal/delivery/http/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoute(app *fiber.App, userHandler *handler.UserHandler, authMiddleware fiber.Handler) {
	public := app.Group("/api")
	public.Post("/register", userHandler.Register)
	public.Post("/login", userHandler.Login)

	protect := app.Group("/api", authMiddleware)
	protect.Get("/user", userHandler.Current)
	protect.Put("/user", userHandler.Update)
}
