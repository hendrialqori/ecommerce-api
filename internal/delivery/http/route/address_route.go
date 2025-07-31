package route

import (
	"internship-mini-project/internal/delivery/http/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterAddressRoute(app *fiber.App, addressHandler *handler.AddressHandler, authMiddleware fiber.Handler) {
	protect := app.Group("/api", authMiddleware)
	protect.Get("/user/alamat", addressHandler.FindAll)
	protect.Get("/user/alamat/:id_alamat", addressHandler.FindById)
	protect.Post("/user/alamat", addressHandler.Create)
	protect.Put("/user/alamat/:id_alamat", addressHandler.Update)
	protect.Delete("/user/alamat/:id_alamat", addressHandler.Delete)
}
