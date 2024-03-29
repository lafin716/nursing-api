package api

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/common/response"
	"nursing_api/internal/domain/plan"
	"nursing_api/pkg/jwt"
)

type PlanHttpApi interface {
	Add(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	GetByDate(ctx *fiber.Ctx) error
	Summary(ctx *fiber.Ctx) error
	Take(ctx *fiber.Ctx) error
	PillToggle(ctx *fiber.Ctx) error
	UpdateMemo(ctx *fiber.Ctx) error
}

type planHttpApi struct {
	service   plan.UseCase
	jwtClient *jwt.JwtClient
}

func NewPlanHttpApi(
	service plan.UseCase,
	jwtClient *jwt.JwtClient,
) PlanHttpApi {
	return &planHttpApi{
		service:   service,
		jwtClient: jwtClient,
	}
}

// @summary 복약 계획 등록
// @description 복약계획을 생성하는 엔드포인트
// @accept json
// @produce json
// @param dto body plan.AddPlanRequest true "복용계획정보"
// @router /plan [post]
// @Security Bearer
func (p planHttpApi) Add(ctx *fiber.Ctx) error {
	req := new(plan.AddPlanRequest)
	err := ctx.BodyParser(req)
	if err != nil {
		return FailParam(err.Error(), ctx)
	}

	errs := validateParameter(req)
	if errs != nil {
		return FailParam(errs, ctx)
	}

	userId, err := getUserId(p.jwtClient, ctx)
	if err != nil {
		return FailAuth(err.Error(), ctx)
	}
	req.UserId = userId
	resp := p.service.Add(req)
	if !resp.Success {
		return Fail(resp.Message, resp.Error, ctx)
	}

	return Ok(nil, ctx)
}

// @summary 복약 계획 삭제
// @description 복약계획을 삭제하는 엔드포인트
// @accept json
// @produce json
// @router /plan/:id [delete]
// @Security Bearer
func (p planHttpApi) Delete(ctx *fiber.Ctx) error {
	req := new(plan.DeletePlanRequest)
	err := ctx.ParamsParser(req)
	if err != nil {
		return FailParam(err.Error(), ctx)
	}

	resp := p.service.Delete(req)
	if !resp.Success {
		return response.New(response.CODE_FAIL_DELETE_PLAN).
			SetMessage(resp.Message).
			SetErrors(resp.Error).
			Error(ctx)
	}

	return response.New(response.CODE_SUCCESS).
		SetMessage(resp.Message).
		Ok(ctx)
}

// @summary 날짜별 복약 계획
// @description 해당 날짜의 복약계획을 조회하는 엔드포인트, 복용상태 및 복용시간을 같이 응답한다.
// @accept json
// @produce json
// @param date query string false "조회날짜 (YYYY-MM-DD), 미입력 시 현재날짜로 세팅"
// @router /plan [get]
// @Security Bearer
func (p planHttpApi) GetByDate(ctx *fiber.Ctx) error {
	req := new(plan.GetByDateRequest)
	parser := ParseRequest(req, QUERY, p.jwtClient, ctx)
	if parser.Error() != nil {
		return parser.Error()
	}

	req.UserId = parser.GetUserId()
	resp := p.service.GetByDate(req)
	if !resp.Success {
		return Fail(resp.Message, resp.Error, ctx)
	}

	return Ok(resp.Data, ctx)
}

// 복용계획개요
func (p planHttpApi) Summary(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

// @summary 복용처리
// @description 복용계획 시간대 전체 복용/미복용처리 (토글로 동작함)
// @produce json
// @param dto body plan.TakeToggleRequest true "복용시간대 정보"
// @router /plan/take [post]
// @Security Bearer
func (p planHttpApi) Take(ctx *fiber.Ctx) error {
	req := new(plan.TakeToggleRequest)
	parser := ParseRequest(req, BODY, p.jwtClient, ctx)
	if parser.Error() != nil {
		return parser.Error()
	}

	req.UserId = parser.GetUserId()
	resp := p.service.Take(req)
	if !resp.Success {
		return Fail(resp.Message, resp.Error.Error(), ctx)
	}

	return OkWithMessage(resp.Message, nil, ctx)
}

// @summary 개별 의약품 복용처리
// @description 복용계획 시간대내 의약품 복용/미복용처리 (토글로 동작함), 기획 상 전체 복용처리를 한 뒤에 실행 가능
// @produce json
// @param dto body plan.PillToggleRequest true "복용계획 의약품 정보"
// @router /plan/take/pill [post]
// @Security Bearer
func (p planHttpApi) PillToggle(ctx *fiber.Ctx) error {
	req := new(plan.PillToggleRequest)
	parser := ParseRequest(req, BODY, p.jwtClient, ctx)
	if parser.Error() != nil {
		return parser.Error()
	}

	req.UserId = parser.GetUserId()
	resp := p.service.PillToggle(req)
	if !resp.Success {
		return Fail(resp.Message, resp.Error.Error(), ctx)
	}

	return OkWithMessage(resp.Message, nil, ctx)
}

// @summary 메모 업데이트
// @description 날짜별 복용계획 시간대의 메모를 업데이트
// @produce json
// @param dto body plan.UpdateMemoRequest true "복용계획 메모 정보"
// @router /plan/memo [post]
// @Security Bearer
func (p planHttpApi) UpdateMemo(ctx *fiber.Ctx) error {
	req := new(plan.UpdateMemoRequest)
	parser := ParseRequest(req, BODY, p.jwtClient, ctx)
	if parser.Error() != nil {
		return parser.Error()
	}

	req.UserId = parser.GetUserId()
	resp := p.service.UpdateMemo(req)
	if !resp.Success {
		return Fail(resp.Message, resp.Error.Error(), ctx)
	}

	return OkWithMessage(resp.Message, nil, ctx)
}
