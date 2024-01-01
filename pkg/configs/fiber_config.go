package configs

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/app/controllers"
	"os"
)

func FiberConfig() fiber.Config {
	appName := os.Getenv("APP_NAME")
	return fiber.Config{
		AppName:      appName,
		ErrorHandler: controllers.GlobalErrorHandler,
	}
}
