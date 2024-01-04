package routes

import (
	"nursing_api/app/controllers"
)

type Routes struct {
	handler *controllers.FiberHandler
}

func NewRoutes(handler *controllers.FiberHandler) *Routes {
	return &Routes{
		handler: handler,
	}
}

func (r *Routes) Resolve() {
	PublicRoutes(r)
	PrivateRoutes(r)
	NotFoundRoute(r)
}
