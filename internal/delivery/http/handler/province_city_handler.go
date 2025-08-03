package handler

import (
	"internship-mini-project/internal/helper"
	"internship-mini-project/internal/model"
	"internship-mini-project/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ProvinceCityHandler struct {
	ProvinceCityUseCase usecase.ProvinceCityUseCase
	Logger              *logrus.Logger
}

func (h *ProvinceCityHandler) FindAllProvince(c *fiber.Ctx) error {
	provinces, err := h.ProvinceCityUseCase.FindAllProvince()
	if err != nil {
		h.Logger.WithError(err).Error(err.Error())
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&model.WebResponse[*[]model.ProvinceResponse]{
		Data:       provinces,
		Message:    "Successfully retrieved data provices",
		StatusCode: fiber.StatusOK,
	})
}

func (h *ProvinceCityHandler) FindAllCityByProvincy(c *fiber.Ctx) error {
	provId, err := helper.ParseToInt(c.Params("prov_id"))
	if err != nil {
		h.Logger.WithError(err).Error(err.Error())
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Prov ID")
	}

	cities, err := h.ProvinceCityUseCase.FindAllCityByProvincy(provId)
	if err != nil {
		h.Logger.WithError(err).Error(err.Error())
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&model.WebResponse[*[]model.CityResponse]{
		Data:       cities,
		Message:    "Successfully retrieved data cities",
		StatusCode: fiber.StatusOK,
	})
}

func (h *ProvinceCityHandler) FindProvinceById(c *fiber.Ctx) error {
	provId, err := helper.ParseToInt(c.Params("prov_id"))
	if err != nil {
		h.Logger.WithError(err).Error(err.Error())
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Prov ID")
	}

	province, err := h.ProvinceCityUseCase.FindProvinceById(provId)
	if err != nil {
		h.Logger.WithError(err).Error(err.Error())
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&model.WebResponse[*model.ProvinceResponse]{
		Data:       province,
		Message:    "Successfully retrieved data provices",
		StatusCode: fiber.StatusOK,
	})
}

func (h *ProvinceCityHandler) FindCityById(c *fiber.Ctx) error {
	cityId, err := helper.ParseToInt(c.Params("city_id"))
	if err != nil {
		h.Logger.WithError(err).Error(err.Error())
		return fiber.NewError(fiber.StatusBadRequest, "Invalid City ID")
	}

	city, err := h.ProvinceCityUseCase.FindCityById(cityId)
	if err != nil {
		h.Logger.WithError(err).Error(err.Error())
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&model.WebResponse[*model.CityResponse]{
		Data:       city,
		Message:    "Successfully retrieved data provices",
		StatusCode: fiber.StatusOK,
	})
}

func NewProvinceCityHandler(provinceCityUseCase usecase.ProvinceCityUseCase, logger *logrus.Logger) *ProvinceCityHandler {
	return &ProvinceCityHandler{
		ProvinceCityUseCase: provinceCityUseCase,
		Logger:              logger,
	}
}
