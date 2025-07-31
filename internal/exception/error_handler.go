package exception

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Errors     any    `json:"errors,omitempty"`
}

// define error here
var (
	// error user domain
	ErrUserAlreadyExist     = fiber.NewError(fiber.StatusBadRequest, "username already exist")
	ErrUserPasswordNotMatch = fiber.NewError(fiber.StatusBadRequest, "password not match")
	ErrUserUnauthorized     = fiber.NewError(fiber.StatusUnauthorized, "User unauthorized")

	//error
	ErrDataAlreadyExists   = fiber.NewError(fiber.StatusBadRequest, "data already exists")
	ErrDataNotFound        = fiber.NewError(fiber.StatusNotFound, "data is not found")
	ErrInternalServerError = fiber.ErrInternalServerError
)

func NewErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		statusCode := fiber.StatusInternalServerError

		var message string
		var errors any

		if e, ok := err.(*fiber.Error); ok {
			statusCode = e.Code
			message = e.Message
		}

		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorMessages := make([]string, 0)
			for _, value := range validationErrors {
				errorMessages = append(errorMessages, fmt.Sprintf(
					"[%s]: '%v' | needs to implements '%s'",
					value.Field(),
					value.Value(),
					value.ActualTag(),
				))
			}

			statusCode = fiber.StatusBadRequest
			errors = errorMessages
		}

		return c.Status(statusCode).JSON(ErrorResponse{
			StatusCode: statusCode,
			Message:    message,
			Errors:     errors,
		})
	}
}
