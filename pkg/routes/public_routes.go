package routes

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/app/controllers"
)

func PublicRoutes(app *fiber.App) {
	route := app.Group("/api/v1")

	route.Post("/user/signup", controllers.UserSignUp)
	route.Post("/user/signin", controllers.UserSignIn)
}
