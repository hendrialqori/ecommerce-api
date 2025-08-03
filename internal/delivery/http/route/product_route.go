package route

import (
	"internship-mini-project/internal/delivery/http/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterProductRoute(app *fiber.App, productHandler *handler.ProductHandler, authMiddleware fiber.Handler) {
	protect := app.Group("/api", authMiddleware)
	protect.Get("/product", productHandler.FindAll)
	protect.Post("/product", productHandler.Create)
	protect.Get("/product/:id_product", productHandler.FindById)
	protect.Put("/product/:id_product", productHandler.Update)
	protect.Delete("/product/:id_product", productHandler.Delete)
}
