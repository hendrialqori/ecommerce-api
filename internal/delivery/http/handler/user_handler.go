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

	err := h.UserUseCase.Register(c.Context(), req)
	if err != nil {
		h.Logger.WithError(err).Error("error user login")
		return err
	}

	return c.
		Status(fiber.StatusOK).
		JSON(&model.WebResponse[any]{
			StatusCode: fiber.StatusOK,
			Message:    "User registered successfully",
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
			StatusCode: fiber.StatusOK,
			Message:    "User logged in successfully",
			Data:       res,
		})
}

func (h *UserHandler) Current(c *fiber.Ctx) error {
	auth := c.Locals("auth").(*model.Auth)

	res, err := h.UserUseCase.Current(c.Context(), auth.Email)
	if err != nil {
		h.Logger.WithError(err).Error("error getting current user")
		return err
	}

	return c.
		Status(fiber.StatusOK).
		JSON(&model.WebResponse[*model.UserResponse]{
			Data:       res,
			Message:    "Current user retrieved successfully",
			StatusCode: fiber.StatusOK,
		})
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	auth := c.Locals("auth").(*model.Auth)

	req := &model.UpdateUserRequest{}
	req.Email = auth.Email

	if err := c.BodyParser(req); err != nil {
		h.Logger.WithError(err).Error("error parsing request body")
		return err
	}

	res, err := h.UserUseCase.Update(c.Context(), req)
	if err != nil {
		h.Logger.WithError(err).Error("error updating user")
		return err
	}

	return c.
		Status(fiber.StatusOK).
		JSON(&model.WebResponse[*model.UserResponse]{
			Data:       res,
			Message:    "User updated successfully",
			StatusCode: fiber.StatusOK,
		})
}
