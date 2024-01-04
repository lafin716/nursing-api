package routes

import (
	"nursing_api/app/controllers"
)

func PrivateRoutes(route *Routes) {
	h := route.handler
	c := h.Container
	app := h.Container.App
	jwtProtector := c.JwtClient.Middleware.JwtProtected()

	app.Group("/api/v1").
		Post("/user/signout", jwtProtector, h.Handle(controllers.UserSignOut))
}
