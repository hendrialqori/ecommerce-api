package route

import (
	"internship-mini-project/internal/delivery/http/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterTokoRoute(app *fiber.App, tokoHandler *handler.TokoHandler, authMiddleware fiber.Handler) {
	protect := app.Group("/api", authMiddleware)
	protect.Get("/toko", tokoHandler.FindAll)
	protect.Get("/toko/my", tokoHandler.Current)
	protect.Get("/toko/:id_toko", tokoHandler.FindById)
	protect.Put("/toko/:id_toko", tokoHandler.Update)
}
