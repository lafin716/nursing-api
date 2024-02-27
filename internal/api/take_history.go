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
}

type takeHistoryHttpApi struct {
	service   takehistory.UseCase
	jwtClient *jwt.JwtClient
}

func NewTakeHistoryHttpApi(
	service takehistory.UseCase,
	jwtClient *jwt.JwtClient,
) TakeHistoryHttpApi {
	return &takeHistoryHttpApi{
		service:   service,
		jwtClient: jwtClient,
	}
}

// @summary 복용내역 목록
// @description 복용내역을 조회하는 엔드포인트
// @accept json
// @produce json
// @param dto query takehistory.GetListRequest true "복용내역 조회 조건 파라미터"
// @router /takehistory [get]
// @Security Bearer
func (t takeHistoryHttpApi) GetList(ctx *fiber.Ctx) error {
	req := new(takehistory.GetListRequest)
	parser := ParseRequest(req, QUERY, t.jwtClient, ctx)
	if parser.Error() != nil {
		return FailParam(parser.Error(), ctx)
	}
	req.UserId = parser.GetUserId()

	resp := t.service.GetList(req)
	if !resp.Success {
		return response.New(response.CODE_NO_DATA).SetErrors(resp.Error).Error(ctx)
	}

	return OkWithMessage(resp.Message, resp.Data, ctx)
}

// @summary 복용내역 상세
// @description 복용내역 상세를 조회하는 엔드포인트
// @accept json
// @produce json
// @router /takehistory/:id [get]
// @Security Bearer
func (t takeHistoryHttpApi) GetDetail(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
