package api

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/common/response"
	"nursing_api/internal/domain/user"
	"nursing_api/pkg/jwt"
)

type UserHttpApi interface {
	Me(ctx *fiber.Ctx) error
	Leave(ctx *fiber.Ctx) error
}

type userHttpApi struct {
	userUseCase user.UserUseCase
	jwtClient   *jwt.JwtClient
}

func NewUserHttpApi(
	userUseCase user.UserUseCase,
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

func (a userHttpApi) Leave(ctx *fiber.Ctx) error {
	claims, err := a.jwtClient.Parser.ExtractTokenMetadata(ctx)
	if err != nil {
		return response.New(response.CODE_INVALID_JWT).
			SetErrors(err).
			Error(ctx)
	}

	userId := claims.UserID
	result, err := a.userUseCase.Leave(userId)
	if err != nil {
		return response.New(response.CODE_FAIL_LEAVE).
			SetErrors(err).
			Error(ctx)
	}

	if !result {
		return response.New(response.CODE_FAIL_LEAVE).
			Error(ctx)
	}

	return response.New(response.CODE_SUCCESS).
		Ok(ctx)
}
