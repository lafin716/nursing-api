package api

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/common/response"
	"nursing_api/internal/domain/auth"
	"nursing_api/pkg/jwt"
)

type AuthHttpApi interface {
	SignIn(ctx *fiber.Ctx) error
	SignUp(ctx *fiber.Ctx) error
	SignOut(ctx *fiber.Ctx) error
}

type authHttpApi struct {
	authUseCase auth.AuthUseCase
	jwtClient   *jwt.JwtClient
}

func NewAuthHttpApi(
	authUseCase auth.AuthUseCase,
	jwtClient *jwt.JwtClient,
) AuthHttpApi {
	return &authHttpApi{
		authUseCase: authUseCase,
		jwtClient:   jwtClient,
	}
}

// @summary 로그인 인증
// @description 회원정보로 인증을 수행하는 엔드포인트. 로그인 성공 시 JWT 토큰 반환.
// @accept json
// @produce json
// @param dto body auth.SignInRequest true "로그인 정보 (이메일, 비밀번호)"
// @router /auth/signin [post]
func (a authHttpApi) SignIn(ctx *fiber.Ctx) error {
	req := new(auth.SignInRequest)
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

// @summary 회원가입
// @description 회원가입을 처리하는 엔드포인트. 회원가입 성공 시 JWT 토큰 반환.
// @accept json
// @produce json
// @param dto body auth.SignUpRequest true "회원가입 정보 (이름, 이메일, 비밀번호)"
// @router /auth/signup [post]
func (a authHttpApi) SignUp(ctx *fiber.Ctx) error {
	req := new(auth.SignUpRequest)
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
