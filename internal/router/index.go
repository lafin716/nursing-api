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
	auth         api.AuthHttpApi
	user         api.UserHttpApi
	medicine     api.MedicineHttpApi
	prescription api.PrescriptionApi
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
	c.RegisterUserRoute()
	c.RegisterAuthRoute()
	c.RegisterMedicineRoute()
	c.RegisterPrescriptionRoute()
}

func NewRouter(
	jwtAuth *middleware.TokenVerifyMiddleware,
	authApi api.AuthHttpApi,
	userApi api.UserHttpApi,
	medicineApi api.MedicineHttpApi,
	prescription api.PrescriptionApi,
) Routable {
	return &container{
		middleware: &middlewares{
			jwtAuth: jwtAuth,
		},
		handler: &handlers{
			authApi,
			userApi,
			medicineApi,
			prescription,
		},
	}
}

func (c *container) AuthMiddleware(handler fiber.Handler) []fiber.Handler {
	mw := c.middleware.jwtAuth.Resolve()
	mw = append(mw, handler)
	return mw
}
