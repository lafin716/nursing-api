package routes

import (
	"nursing_api/app/controllers"
)

func PublicRoutes(route *Routes) {
	h := route.handler
	app := route.handler.Container.App

	app.Get("/", h.Handle(controllers.Hello))
	app.Group("/api/v1").
		Post("/user/signup", h.Handle(controllers.UserSignUp)).
		Post("/user/signin", h.Handle(controllers.UserSignIn))
}
