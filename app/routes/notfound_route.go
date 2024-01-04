package routes

import (
	"nursing_api/app/controllers"
)

func NotFoundRoute(route *Routes) {
	h := route.handler
	app := h.Container.App
	app.Use(h.Handle(controllers.NotFound))
}
