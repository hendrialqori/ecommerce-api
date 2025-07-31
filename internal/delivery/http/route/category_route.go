package route

import (
	"internship-mini-project/internal/delivery/http/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterCategoryRoute(app *fiber.App, categoryHandler *handler.CategoryHandler, authMiddleware fiber.Handler) {
	protect := app.Group("/api", authMiddleware)
	protect.Get("/category", categoryHandler.FindAll)
	protect.Get("/category/:id_category", categoryHandler.FindById)
	protect.Post("/category", categoryHandler.Create)
	protect.Put("/category/:id_category", categoryHandler.Update)
	protect.Delete("/category/:id_category", categoryHandler.Delete)
}
