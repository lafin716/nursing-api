package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"nursing_api/internal/api"
	"nursing_api/pkg/jwt"
)

var Set = wire.NewSet(
	NewRouter,
)

type Routable interface {
	Init(router *fiber.Router, jwtMiddleware *jwt.JwtMiddleware)
}

type apiSet struct {
	authApi     api.AuthHttpApi
	userApi     api.UserHttpApi
	medicineApi api.MedicineHttpApi
}

func NewRouter(
	authApi api.AuthHttpApi,
	userApi api.UserHttpApi,
	medicineApi api.MedicineHttpApi,
) Routable {
	return &apiSet{
		authApi,
		userApi,
		medicineApi,
	}
}

func (r *apiSet) Init(
	router *fiber.Router,
	jwtMiddleware *jwt.JwtMiddleware,
) {
	NewAuthRouter(r).Init(router, jwtMiddleware)
	NewUserRouter(r).Init(router, jwtMiddleware)
	NewMedicineRouter(r).Init(router, jwtMiddleware)
}
