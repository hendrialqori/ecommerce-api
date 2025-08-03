package route

import (
	"internship-mini-project/internal/delivery/http/handler"

	"github.com/gofiber/fiber/v2"
)

func NewRegisterProvinceCityRoute(app *fiber.App, provinceCityHandler *handler.ProvinceCityHandler, authMiddleware fiber.Handler) {
	protect := app.Group("/api", authMiddleware)
	protect.Get("/provcity/listprovincies", provinceCityHandler.FindAllProvince)
	protect.Get("/provcity/listcities/:prov_id", provinceCityHandler.FindAllCityByProvincy)
	protect.Get("/provcity/detailprovince/:prov_id", provinceCityHandler.FindProvinceById)
	protect.Get("/provcity/detailcity/:city_id", provinceCityHandler.FindCityById)
}
