package route

import (
	"internship-mini-project/internal/delivery/http/handler"

	"github.com/gofiber/fiber/v2"
)

func NewRegisterProvinceCityRoute(app *fiber.App, provinceCityHandler *handler.ProvinceCityHandler, authMiddleware fiber.Handler) {
	public := app.Group("/api")
	public.Get("/provcity/listprovincies", provinceCityHandler.FindAllProvince)
	public.Get("/provcity/listcities/:prov_id", provinceCityHandler.FindAllCityByProvincy)
	public.Get("/provcity/detailprovince/:prov_id", provinceCityHandler.FindProvinceById)
	public.Get("/provcity/detailcity/:city_id", provinceCityHandler.FindCityById)
}
