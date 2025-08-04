package route

import (
	"internship-mini-project/internal/delivery/http/handler"

	"github.com/gofiber/fiber/v2"
)

func NewRegisterTrxRoute(app *fiber.App, trxHandler *handler.TrxHandler, authMiddleware fiber.Handler) {
	protect := app.Group("/api", authMiddleware)
	protect.Get("/trx", trxHandler.FindAll)
	protect.Get("/trx/:id_trx", trxHandler.FindById)
	protect.Post("/trx", trxHandler.Create)
}
