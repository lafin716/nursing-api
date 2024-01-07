package api

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/common/response"
	"nursing_api/internal/domain/auth/dto"
	"nursing_api/internal/domain/auth/usecase"
	"nursing_api/pkg/jwt"
)

type AuthHttpApi interface {
	SignIn(ctx *fiber.Ctx) error
	SignUp(ctx *fiber.Ctx) error
	SignOut(ctx *fiber.Ctx) error
}

type authHttpApi struct {
	authUseCase usecase.AuthUseCase
	jwtClient   *jwt.JwtClient
}

func NewAuthHttpApi(
	authUseCase usecase.AuthUseCase,
	jwtClient *jwt.JwtClient,
) AuthHttpApi {
	return &authHttpApi{
		authUseCase: authUseCase,
		jwtClient:   jwtClient,
	}
}

func (a authHttpApi) SignIn(ctx *fiber.Ctx) error {
	req := new(dto.SignInRequest)
	err := ctx.BodyParser(req)
	if err != nil {
		return response.New(response.CODE_INVALID_PARAM).
			SetErrors(err).
			Error(ctx)
	}

	result := a.authUseCase.SignIn(req)
	if !result.Success {
		return response.New(response.CODE_FAIL_SIGN_UP).
			SetMessage(result.Message).
			SetErrors(result.Error).
			Error(ctx)
	}

	return response.New(response.CODE_SUCCESS).
		SetData(result.Token).
		Ok(ctx)
}

func (a authHttpApi) SignUp(ctx *fiber.Ctx) error {
	req := new(dto.SignUpRequest)
	err := ctx.BodyParser(req)
	if err != nil {
		return response.New(response.CODE_INVALID_PARAM).
			SetErrors(err).
			Error(ctx)
	}

	result := a.authUseCase.SignUp(req)
	if !result.Success {
		return response.New(response.CODE_FAIL_SIGN_UP).
			SetMessage(result.Message).
			SetErrors(result.Error).
			Error(ctx)
	}

	return response.New(response.CODE_SUCCESS).
		SetData(result.Token).
		Ok(ctx)
}

func (a authHttpApi) SignOut(ctx *fiber.Ctx) error {
	_, err := a.jwtClient.Parser.ExtractTokenMetadata(ctx)
	if err != nil {
		return response.New(response.CODE_INVALID_JWT).
			SetErrors(err).
			Error(ctx)
	}

	return response.New(response.CODE_SUCCESS).
		SetMessage("로그아웃 되었습니다.").
		Ok(ctx)
}
