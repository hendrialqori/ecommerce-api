package infrastructure

import (
	"time"

	"internship-mini-project/internal/exception"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

func NewFiber(config *viper.Viper) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      config.GetString("APP_NAME"),
		ErrorHandler: exception.NewErrorHandler(),
		Prefork:      config.GetBool("APP_PREFORK"),
		WriteTimeout: config.GetDuration("APP_TIMEOUT") * time.Second,
		ReadTimeout:  config.GetDuration("APP_TIMEOUT") * time.Second,
	})
	app.Use(recover.New())

	return app
}
