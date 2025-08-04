package handler

import (
	"internship-mini-project/internal/helper"
	"internship-mini-project/internal/model"
	"internship-mini-project/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type TrxHandler struct {
	TrxUseCase usecase.TrxUseCase
	Logger     *logrus.Logger
}

func (h *TrxHandler) FindAll(c *fiber.Ctx) error {
	transaction, err := h.TrxUseCase.FindAll(c.Context())
	if err != nil {
		h.Logger.WithError(err).Error("Failed to find all transactions")
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&model.WebResponse[[]*model.TrxResponse]{
		Data:       transaction,
		Message:    "Successfully retrieved all transaction",
		StatusCode: fiber.StatusOK,
	})
}

func (h *TrxHandler) FindById(c *fiber.Ctx) error {
	id, err := helper.ParseToInt(c.Params("id_trx"))
	if err != nil {
		h.Logger.WithError(err).Error("Invalid Transation ID")
		return err
	}

	trx, err := h.TrxUseCase.FindById(c.Context(), id)
	if err != nil {
		h.Logger.WithError(err).Error("Failed to transation trx by ID")
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&model.WebResponse[*model.TrxResponse]{
		Data:       trx,
		Message:    "Successfully retrieved transations by ID",
		StatusCode: fiber.StatusOK,
	})
}

func (h *TrxHandler) Create(c *fiber.Ctx) error {
	auth := c.Locals("auth").(*model.Auth)

	req := &model.CreateTrxRequest{}
	req.IDUser = auth.ID

	if err := c.BodyParser(req); err != nil {
		h.Logger.WithError(err).Error("error parsing request body")
		return err
	}

	err := h.TrxUseCase.Create(c.Context(), req)
	if err != nil {
		h.Logger.WithError(err).Error(err.Error())
		return err
	}

	return c.
		Status(fiber.StatusCreated).
		JSON(&model.WebResponse[*model.UserResponse]{
			StatusCode: fiber.StatusCreated,
			Message:    "Transaction created successfully",
		})
}

func NewTrxHandler(TrxUseCase usecase.TrxUseCase, Logger *logrus.Logger) *TrxHandler {
	return &TrxHandler{
		TrxUseCase: TrxUseCase,
		Logger:     Logger,
	}
}
