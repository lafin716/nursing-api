package api

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/common/response"
	takehistory "nursing_api/internal/domain/take_history"
	"nursing_api/pkg/jwt"
)

type TakeHistoryHttpApi interface {
	GetList(ctx *fiber.Ctx) error
	GetDetail(ctx *fiber.Ctx) error
	TakePlan(ctx *fiber.Ctx) error
	TakePill(ctx *fiber.Ctx) error
}

type takeHistoryHttpApi struct {
	service   takehistory.TakeHistoryUseCase
	jwtClient *jwt.JwtClient
}

func NewTakeHistoryHttpApi(
	service takehistory.TakeHistoryUseCase,
	jwtClient *jwt.JwtClient,
) TakeHistoryHttpApi {
	return &takeHistoryHttpApi{
		service:   service,
		jwtClient: jwtClient,
	}
}

func (t takeHistoryHttpApi) GetList(ctx *fiber.Ctx) error {
	req := new(takehistory.GetListRequest)
	err := ctx.QueryParser(req)
	if err != nil {
		return response.New(response.CODE_INVALID_PARAM).SetErrors(err).Error(ctx)
	}

	userId, err := getUserId(t.jwtClient, ctx)
	if err != nil {
		return err
	}
	req.UserId = userId

	resp := t.service.GetList(req)
	if !resp.Success {
		return response.New(response.CODE_NO_DATA).SetErrors(err).Error(ctx)
	}

	return response.New(response.CODE_SUCCESS).
		SetData(resp.Data).
		SetMessage(resp.Message).
		Ok(ctx)
}

func (t takeHistoryHttpApi) GetDetail(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (t takeHistoryHttpApi) TakePlan(ctx *fiber.Ctx) error {
	req := new(takehistory.TakePlanRequest)
	err := ctx.BodyParser(req)
	if err != nil {
		return response.New(response.CODE_INVALID_PARAM).SetErrors(err).Error(ctx)
	}

	userId, err := getUserId(t.jwtClient, ctx)
	if err != nil {
		return err
	}
	req.UserId = userId

	resp := t.service.TakePlanToggle(req)
	if !resp.Success {
		return response.New(response.CODE_NO_DATA).SetErrors(err).Error(ctx)
	}

	return response.New(response.CODE_SUCCESS).
		SetMessage(resp.Message).
		Ok(ctx)
}

func (t takeHistoryHttpApi) TakePill(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
