package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"nursing_api/internal/api"
	"nursing_api/internal/middleware"
)

var Set = wire.NewSet(
	NewRouter,
)

type Routable interface {
	Init(router *fiber.Router)
}

type middlewares struct {
	jwtAuth *middleware.TokenVerifyMiddleware
}

type handlers struct {
	main         api.MainHttpApi
	auth         api.AuthHttpApi
	user         api.UserHttpApi
	medicine     api.MedicineHttpApi
	prescription api.PrescriptionApi
	plan         api.PlanHttpApi
	timezone     api.TimeZoneApi
}

type container struct {
	router     fiber.Router
	middleware *middlewares
	handler    *handlers
}

func (c *container) Init(
	router *fiber.Router,
) {
	c.router = *router
	c.RegisterMainRoute()
	c.RegisterUserRoute()
	c.RegisterAuthRoute()
	c.RegisterMedicineRoute()
	c.RegisterPrescriptionRoute()
	c.RegisterPlanRoute()
	c.RegisterTimeZoneRoute()
}

func NewRouter(
	jwtAuth *middleware.TokenVerifyMiddleware,
	main api.MainHttpApi,
	auth api.AuthHttpApi,
	user api.UserHttpApi,
	medicine api.MedicineHttpApi,
	prescription api.PrescriptionApi,
	plan api.PlanHttpApi,
	timezone api.TimeZoneApi,
) Routable {
	return &container{
		middleware: &middlewares{
			jwtAuth: jwtAuth,
		},
		handler: &handlers{
			main:         main,
			auth:         auth,
			user:         user,
			medicine:     medicine,
			prescription: prescription,
			plan:         plan,
			timezone:     timezone,
		},
	}
}

func (c *container) AuthMiddleware(handler fiber.Handler) []fiber.Handler {
	mw := c.middleware.jwtAuth.Resolve()
	mw = append(mw, handler)
	return mw
}
