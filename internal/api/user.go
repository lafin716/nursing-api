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
	CheckEmail(ctx *fiber.Ctx) error
}

type userHttpApi struct {
	userUseCase user.UseCase
	jwtClient   *jwt.JwtClient
}

func NewUserHttpApi(
	userUseCase user.UseCase,
	jwtClient *jwt.JwtClient,
) UserHttpApi {
	return &userHttpApi{
		userUseCase: userUseCase,
		jwtClient:   jwtClient,
	}
}

// @summary 내 정보 조회
// @description 회원정보를 조회하는 엔드포인트
// @accept json
// @produce json
// @router /user/me [get]
// @Security Bearer
func (a userHttpApi) Me(ctx *fiber.Ctx) error {
	claims, err := a.jwtClient.Parser.ExtractTokenMetadata(ctx)
	if err != nil {
		return response.New(response.CODE_INVALID_JWT).
			SetErrors(err).
			Error(ctx)
	}

	userId := claims.UserID
	result := a.userUseCase.GetUser(userId)
	return ResolveResponse(result, ctx)
}

// @summary 회원탈퇴
// @description 회원탈퇴 엔드포인트
// @accept json
// @produce json
// @router /user/leave [delete]
// @Security Bearer
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

// @summary 이메일 중복확인
// @description 이메일 중복확인 엔드포인트
// @accept json
// @produce json
// @router /user/check-email [post]
func (a userHttpApi) CheckEmail(ctx *fiber.Ctx) error {
	req := new(user.CheckEmailRequest)
	err := ctx.BodyParser(req)
	if err != nil {
		return response.New(response.CODE_INVALID_PARAM).
			SetErrors(err).
			Error(ctx)
	}

	errs := validateParameter(req)
	if errs != nil {
		return response.New(response.CODE_INVALID_PARAM).
			SetErrors(errs).
			Error(ctx)
	}

	resp := a.userUseCase.CheckDuplicatedEmail(req)
	return ResolveResponse(resp, ctx)
}
