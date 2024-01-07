package api

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/common/response"
	"nursing_api/internal/domain/auth/dto"
	"nursing_api/internal/domain/auth/usecase"
)

type AuthHttpApi interface {
	SignIn(ctx *fiber.Ctx) error
	SignUp(ctx *fiber.Ctx) error
	SignOut(ctx *fiber.Ctx) error
}

type authHttpApi struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthHttpApi(authUseCase usecase.AuthUseCase) AuthHttpApi {
	return &authHttpApi{
		authUseCase: authUseCase,
	}
}

func (a authHttpApi) SignIn(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
}
