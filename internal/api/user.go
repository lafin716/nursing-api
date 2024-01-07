package api

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/common/response"
	"nursing_api/internal/domain/user/usecase"
	"nursing_api/pkg/jwt"
)

type UserHttpApi interface {
	Me(ctx *fiber.Ctx) error
}

type userHttpApi struct {
	userUseCase usecase.UserUseCase
	jwtClient   *jwt.JwtClient
}

func NewUserHttpApi(
	userUseCase usecase.UserUseCase,
	jwtClient *jwt.JwtClient,
) UserHttpApi {
	return &userHttpApi{
		userUseCase: userUseCase,
		jwtClient:   jwtClient,
	}
}

func (h *userHttpApi) Me(ctx *fiber.Ctx) error {
	claims, err := h.jwtClient.Parser.ExtractTokenMetadata(ctx)
	if err != nil {
		return response.New(response.CODE_INVALID_JWT).
			SetErrors(err).
			Error(ctx)
	}

	userId := claims.UserID
	result := h.userUseCase.GetUser(userId)
	if !result.Success {
		return response.New(response.CODE_USER_NOTFOUND).
			SetMessage(result.Message).
			SetErrors(result.Error).
			Error(ctx)
	}

	return response.New(response.CODE_SUCCESS).
		SetData(result.User).
		Ok(ctx)
}
