package route

import (
	"internship-mini-project/internal/delivery/http/handler"

	"github.com/gofiber/fiber/v2"
)

func NewRegisterProductPhotoRoute(app *fiber.App, handler *handler.ProductPhotoHandler, authMiddleware fiber.Handler) {
	protect := app.Group("/api", authMiddleware)
	protect.Post("/product/foto", handler.Create)
	protect.Delete("/product/foto/:id_foto", handler.Delete)
}
