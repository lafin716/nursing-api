package routes

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/app/controllers"
	"nursing_api/pkg/middlewares"
)

func PrivateRoutes(app *fiber.App) {
	route := app.Group("/api/v1")

	route.Post("/user/signout", middlewares.JwtProtected(), controllers.UserSignOut)
}
