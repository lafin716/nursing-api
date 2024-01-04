package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type FiberClient struct {
	app    *fiber.App
	config *FiberConfig
}

func NewFiberClient(config *FiberConfig) *FiberClient {
	app := fiber.New(fiberConfig(config))
	app.Use(
		cors.New(),
		logger.New(),
	)

	return &FiberClient{
		app:    app,
		config: config,
	}
}

func fiberConfig(config *FiberConfig) fiber.Config {
	appName := config.AppName
	return fiber.Config{
		AppName:      appName,
		ErrorHandler: config.ErrorHandler,
	}
}

func (client *FiberClient) GetApp() *fiber.App {
	return client.app
}
