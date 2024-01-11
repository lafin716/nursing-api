package router

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/api"
	"nursing_api/pkg/jwt"
)

type userRouter struct {
	userApi api.UserHttpApi
}

func NewUserRouter(apiSet *apiSet) Routable {
	return &userRouter{
		userApi: apiSet.userApi,
	}
}

func (r *userRouter) Init(
	router *fiber.Router,
	jwtMiddleware *jwt.JwtMiddleware,
) {
	userRouter := (*router).Group("/user")
	{
		userRouter.Get("/me", jwtMiddleware.JwtProtected(), r.userApi.Me)
	}
}
