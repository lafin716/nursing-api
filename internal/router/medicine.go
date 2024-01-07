package router

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/api"
	"nursing_api/pkg/jwt"
)

type MedicineRouter interface {
	Init(router *fiber.Router, jwtMiddleware *jwt.JwtMiddleware)
}

type medicineRouter struct {
	medicineApi api.MedicineHttpApi
}

func NewMedicineRouter(medicineApi api.MedicineHttpApi) MedicineRouter {
	return &medicineRouter{
		medicineApi: medicineApi,
	}
}

func (r *medicineRouter) Init(
	router *fiber.Router,
	jwtMiddleware *jwt.JwtMiddleware,
) {
	userRouter := (*router).Group("/medicine")
	{
		userRouter.Get("/search/:pillName", jwtMiddleware.JwtProtected(), r.medicineApi.Search)
	}
}
