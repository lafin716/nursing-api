package router

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/api"
	"nursing_api/pkg/jwt"
)

type AuthRouter interface {
	Init(router *fiber.Router, jwtMiddleware *jwt.JwtMiddleware)
}

type authRouter struct {
	authApi api.AuthHttpApi
}

func NewAuthRouter(authApi api.AuthHttpApi) AuthRouter {
	return &authRouter{
		authApi: authApi,
	}
}

func (r *authRouter) Init(
	router *fiber.Router,
	jwtMiddleware *jwt.JwtMiddleware,
) {
	userRouter := (*router).Group("/user")
	{
		userRouter.Post("/signup", r.authApi.SignUp)
		userRouter.Post("/signin", r.authApi.SignIn)
		userRouter.Post("/signout", jwtMiddleware.JwtProtected(), r.authApi.SignOut)
	}
}
