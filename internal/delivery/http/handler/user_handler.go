package handler

import (
	"internship-mini-project/internal/model"
	"internship-mini-project/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	UserUseCase usecase.UserUseCase
	Logger      *logrus.Logger
}

func NewUserHandler(userUseCase usecase.UserUseCase, logger *logrus.Logger) *UserHandler {
	return &UserHandler{
		UserUseCase: userUseCase,
		Logger:      logger,
	}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	req := &model.RegisterUserRequest{}
	if err := c.BodyParser(req); err != nil {
		h.Logger.WithError(err).Error("error parsing request body")
		return err
	}

	res, err := h.UserUseCase.Register(c.Context(), req)

	if err != nil {
		h.Logger.WithError(err).Error("error user register")
		return err
	}

	return c.
		Status(fiber.StatusCreated).
		JSON(&model.WebResponse[*model.UserResponse]{
			Data: res,
		})
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	req := &model.LoginUserRequest{}
	if err := c.BodyParser(req); err != nil {
		h.Logger.WithError(err).Error("error parsing request body")
		return err
	}

	res, err := h.UserUseCase.Login(c.Context(), req)
	if err != nil {
		h.Logger.WithError(err).Error("error user login")
		return err
	}

	return c.
		Status(fiber.StatusOK).
		JSON(&model.WebResponse[*model.TokenResponse]{
			Data: res,
		})
}
