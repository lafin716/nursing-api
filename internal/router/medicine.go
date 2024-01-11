package router

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/api"
	"nursing_api/pkg/jwt"
)

type medicineRouter struct {
	medicineApi api.MedicineHttpApi
}

func NewMedicineRouter(apiSet *apiSet) Routable {
	return &medicineRouter{
		medicineApi: apiSet.medicineApi,
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
