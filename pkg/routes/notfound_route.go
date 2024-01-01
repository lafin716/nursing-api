package routes

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/app/controllers"
)

func NotFoundRoute(app *fiber.App) {
	app.Use(controllers.NotFound)
}
