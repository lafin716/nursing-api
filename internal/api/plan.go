package api

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/common/response"
	"nursing_api/internal/domain/plan"
	"nursing_api/pkg/jwt"
)

type PlanHttpApi interface {
	Today(ctx *fiber.Ctx) error
	Summary(ctx *fiber.Ctx) error
	TakePlan(ctx *fiber.Ctx) error
	TakePill(ctx *fiber.Ctx) error
	UpdateMemo(ctx *fiber.Ctx) error
}

type planHttpApi struct {
	service   plan.PlanUseCase
	jwtClient *jwt.JwtClient
}

// @summary 날짜별 복약 계획
// @description 해당 날짜의 복약계획을 조회하는 엔드포인트, 복용상태 및 복용시간을 같이 응답한다.
// @accept json
// @produce json
// @param current_date query string true "조회날짜 (YYYY-MM-DD)"
// @router /plan [get]
func (p planHttpApi) Today(ctx *fiber.Ctx) error {
	req := new(plan.GetByDateRequest)
	err := ctx.QueryParser(req)
	if err != nil {
		return response.New(response.CODE_INVALID_PARAM).Error(ctx)
	}

	errs := validateParameter(req)
	if errs != nil {
		return response.New(response.CODE_INVALID_PARAM).SetErrors(errs).Error(ctx)
	}

	userId, err := getUserId(p.jwtClient, ctx)
	if err != nil {
		return response.New(response.CODE_INVALID_JWT).SetErrors(errs).Error(ctx)
	}

	req.UserId = userId
	resp := p.service.GetByDate(req)
	if !resp.Success {
		return response.New(response.CODE_ERROR).SetMessage(resp.Message).SetErrors(resp.Error).Error(ctx)
	}

	return response.New(response.CODE_SUCCESS).SetData(resp.Data).Ok(ctx)
}

func (p planHttpApi) Summary(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (p planHttpApi) TakePlan(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (p planHttpApi) TakePill(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (p planHttpApi) UpdateMemo(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func NewPlanHttpApi(
	service plan.PlanUseCase,
	jwtClient *jwt.JwtClient,
) PlanHttpApi {
	return &planHttpApi{
		service:   service,
		jwtClient: jwtClient,
	}
}
