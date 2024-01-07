package router

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/api"
	"nursing_api/pkg/jwt"
)

type UserRouter interface {
	Init(router *fiber.Router, jwtMiddleware *jwt.JwtMiddleware)
}

type userRouter struct {
	userApi api.UserHttpApi
}

func NewUserRouter(userApi api.UserHttpApi) UserRouter {
	return &userRouter{
		userApi: userApi,
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
