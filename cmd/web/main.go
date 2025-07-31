package main

import (
	"fmt"
	"internship-mini-project/config"
	"internship-mini-project/internal/delivery/http/handler"
	"internship-mini-project/internal/delivery/http/middleware"
	"internship-mini-project/internal/delivery/http/route"
	"internship-mini-project/internal/domain/registry"
	"internship-mini-project/internal/infrastructure"
	"internship-mini-project/internal/repository"
	"internship-mini-project/internal/usecase"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := config.New()

	app := infrastructure.NewFiber(config)
	port := config.Get("APP_PORT")

	db := infrastructure.NewGorm(config)
	logger := infrastructure.NewLogger(config)
	validate := infrastructure.NewValidator(config)

	for _, model := range registry.RegistryModels() {
		if err := db.AutoMigrate(model.Model); err != nil {
			panic(fmt.Errorf("error migrating model %T : %+v", model.Model, err.Error()))
		}
	}

	var (
		tokoRepo    = repository.NewTokoRepository(db)
		tokoUseCase = usecase.NewTokoUseCase(tokoRepo, logger, validate, config)
		tokoHandler = handler.NewTokoHandler(tokoUseCase, logger, config)

		userRepo    = repository.NewUserRepository(db)
		userUseCase = usecase.NewUserUseCase(userRepo, tokoRepo, logger, validate, config)
		userHandler = handler.NewUserHandler(userUseCase, logger)

		categoryRepo    = repository.NewCategoryRepository(db)
		categoryUseCase = usecase.NewCategoryUseCase(categoryRepo, logger, validate)
		categoryHandler = handler.NewCategoryHandler(categoryUseCase, logger)

		addressRepo    = repository.NewAddressRepository(db)
		addressUseCase = usecase.NewAddressUseCase(addressRepo, logger, validate)
		addressHandler = handler.NewAddressHandler(addressUseCase, logger)
	)

	app.Get("/api/ping!", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Welcome to the Internship Mini Project API",
		})
	})

	auth := middleware.NewAuth(logger, config)
	route.RegisterUserRoute(app, userHandler, auth)
	route.RegisterTokoRoute(app, tokoHandler, auth)
	route.RegisterCategoryRoute(app, categoryHandler, auth)
	route.RegisterAddressRoute(app, addressHandler, auth)

	go func() {
		if err := app.Listen(fmt.Sprintf(":%v", port)); err != nil {
			panic(fmt.Errorf("error running app : %+v", err.Error()))
		}
	}()

	ch := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	<-ch // This blocks the main thread until an interrupt is received

	// Your cleanup tasks go here
	_ = app.Shutdown()

	fmt.Println("App was successful shutdown.")
}
