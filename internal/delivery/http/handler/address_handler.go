package handler

import (
	"internship-mini-project/internal/model"
	"internship-mini-project/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AddressHandler struct {
	AdressUseCase usecase.AddressUsecase
	Logger        *logrus.Logger
}

func NewAddressHandler(addressUseCase usecase.AddressUsecase, logger *logrus.Logger) *AddressHandler {
	return &AddressHandler{
		AdressUseCase: addressUseCase,
		Logger:        logger,
	}
}

func (a *AddressHandler) FindAll(c *fiber.Ctx) error {

	auth := c.Locals("auth").(*model.Auth)

	addresses, err := a.AdressUseCase.FindAll(c.Context(), int(auth.ID))
	if err != nil {
		a.Logger.WithError(err).Error("Failed to find all addresses")
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&model.WebResponse[[]*model.AddressResponse]{
		Data:       addresses,
		Message:    "Successfully retrieved all addresses",
		StatusCode: fiber.StatusOK,
	})
}

func (a *AddressHandler) FindById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id_alamat"))
	if err != nil || id <= 0 {
		a.Logger.WithError(err).Error("Invalid address ID")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid address ID")
	}

	address, err := a.AdressUseCase.FindById(c.Context(), id)
	if err != nil {
		a.Logger.WithError(err).Error("Failed to find address by ID")
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&model.WebResponse[*model.AddressResponse]{
		Data:       address,
		Message:    "Address found successfully",
		StatusCode: fiber.StatusOK,
	})
}

func (a *AddressHandler) Create(c *fiber.Ctx) error {
	req := &model.CreateAddressRequest{}
	if err := c.BodyParser(req); err != nil {
		a.Logger.WithError(err).Error("Error parsing request body")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	auth := c.Locals("auth").(*model.Auth)
	req.IDUser = auth.ID
	address, err := a.AdressUseCase.Create(c.Context(), req)
	if err != nil {
		a.Logger.WithError(err).Error("Failed to create address")
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(&model.WebResponse[*model.AddressResponse]{
		Data:       address,
		Message:    "Address created successfully",
		StatusCode: fiber.StatusCreated,
	})
}

func (a *AddressHandler) Update(c *fiber.Ctx) error {
	req := &model.UpdateAddressRequest{}
	if err := c.BodyParser(req); err != nil {
		a.Logger.WithError(err).Error("Error parsing request body")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	id, err := strconv.Atoi(c.Params("id_alamat"))
	if err != nil || id <= 0 {
		a.Logger.WithError(err).Error("Invalid address ID")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid address ID")
	}

	req.ID = uint(id)
	req.IDUser = c.Locals("auth").(*model.Auth).ID

	address, err := a.AdressUseCase.Update(c.Context(), req)
	if err != nil {
		a.Logger.WithError(err).Error("Failed to update address")
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&model.WebResponse[*model.AddressResponse]{
		Data:       address,
		Message:    "Address updated successfully",
		StatusCode: fiber.StatusOK,
	})
}

func (a *AddressHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id_alamat"))
	if err != nil || id <= 0 {
		a.Logger.WithError(err).Error("Invalid address ID")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid address ID")
	}

	if err := a.AdressUseCase.Delete(c.Context(), id); err != nil {
		a.Logger.WithError(err).Error("Failed to delete address")
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&model.WebResponse[any]{
		Message:    "Address deleted successfully",
		StatusCode: fiber.StatusOK,
	})
}
