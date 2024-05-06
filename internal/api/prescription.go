package api

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/common/response"
	"nursing_api/internal/domain/prescription"
	"nursing_api/pkg/jwt"
)

type PrescriptionApi interface {
	Register(ctx *fiber.Ctx) error
	GetList(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error

	GetItemList(ctx *fiber.Ctx) error
	GetItemById(ctx *fiber.Ctx) error
	AddItem(ctx *fiber.Ctx) error
	UpdateItem(ctx *fiber.Ctx) error
	DeleteItem(ctx *fiber.Ctx) error
}

type prescriptionApi struct {
	service   prescription.UseCase
	jwtClient *jwt.JwtClient
}

func NewPrescriptionApi(
	service prescription.UseCase,
	jwtClient *jwt.JwtClient,
) PrescriptionApi {
	return &prescriptionApi{
		service:   service,
		jwtClient: jwtClient,
	}
}

// @summary 처방전 리스트 조회
// @description 해당 날짜의 처방전 목록을 조회하는 엔드포인트
// @produce json
// @param date query string true "조회날짜 (YYYY-MM-DD)"
// @param limit query int true "조회갯수"
// @router /prescription [get]
// @Security Bearer
func (a *prescriptionApi) GetList(ctx *fiber.Ctx) error {
	req := new(prescription.GetListRequest)
	parser := ParseRequest(req, QUERY, a.jwtClient, ctx)
	if parser.Error() != nil {
		return parser.Error()
	}

	req.UserId = parser.GetUserId()
	resp := a.service.GetList(req)
	if !resp.Success {
		return Fail(resp.Message, resp.Error.Error(), ctx)
	}

	return OkWithMessage(resp.Message, resp.Data, ctx)
}

// @summary 처방전 상세 조회
// @description 처방전 상세정보를 조회하는 엔드포인트
// @produce json
// @param id path string true "처방전 고유번호"
// @router /prescription/:id [get]
// @Security Bearer
func (a *prescriptionApi) GetById(ctx *fiber.Ctx) error {
	req := new(prescription.GetByIdRequest)
	parser := ParseRequest(req, PARAM, a.jwtClient, ctx)
	if parser.Error() != nil {
		return parser.Error()
	}

	req.UserId = parser.GetUserId()
	resp := a.service.GetById(req)
	if !resp.Success {
		return Fail(resp.Message, resp.Error.Error(), ctx)
	}

	return OkWithMessage(resp.Message, resp.Data, ctx)
}

// @summary 처방전 등록
// @description [Deprecated:복용계획에서 등록하는 프로세스로 변경됨] 처방전을 등록하는 엔드포인트
// @produce json
// @param dto body prescription.RegisterRequest true "처방전 등록정보"
// @router /prescription [post]
// @Security Bearer
func (a *prescriptionApi) Register(ctx *fiber.Ctx) error {
	req := new(prescription.RegisterRequest)
	parser := ParseRequest(req, BODY, a.jwtClient, ctx)
	if parser.Error() != nil {
		return parser.Error()
	}
	req.UserId = parser.GetUserId()

	resp := a.service.Register(req)
	if !resp.Success {
		return Fail(resp.Message, resp.Error.Error(), ctx)
	}

	return OkWithMessage(resp.Message, nil, ctx)
}

// @summary 처방전 업데이트
// @description 처방전을 업데이트하는 엔드포인트
// @produce json
// @param dto body prescription.UpdateRequest true "처방전 정보"
// @router /prescription [patch]
// @Security Bearer
func (a *prescriptionApi) Update(ctx *fiber.Ctx) error {
	req := new(prescription.UpdateRequest)
	parser := ParseRequest(req, BODY, a.jwtClient, ctx)
	if parser.Error() != nil {
		return parser.Error()
	}

	resp := a.service.Update(req)
	if !resp.Success {
		return Fail(resp.Message, resp.Error.Error(), ctx)
	}

	return OkWithMessage(resp.Message, nil, ctx)
}

// @summary 처방전 삭제
// @description 처방전을 삭제하는 엔드포인트
// @produce json
// @param id path string true "처방전 고유번호"
// @router /prescription/:id [delete]
// @Security Bearer
func (a *prescriptionApi) Delete(ctx *fiber.Ctx) error {
	req := new(prescription.DeleteRequest)
	parser := ParseRequest(req, PARAM, a.jwtClient, ctx)
	if parser.Error() != nil {
		return parser.Error()
	}

	resp := a.service.Delete(req)
	if !resp.Success {
		return Fail(resp.Message, resp.Error.Error(), ctx)
	}

	return OkWithMessage(resp.Message, nil, ctx)
}

// @summary 처방전 의약품 목록 조회
// @description 처방전 상세에서 의약품 목록을 조회하는 엔드포인트
// @produce json
// @param date query string true "조회날짜 (YYYY-MM-DD)"
// @router /prescription/items [get]
// @Security Bearer
func (a *prescriptionApi) GetItemList(ctx *fiber.Ctx) error {
	req := new(prescription.GetItemListRequest)
	parser := ParseRequest(req, QUERY, a.jwtClient, ctx)
	if parser.Error() != nil {
		return parser.Error()
	}
	req.UserId = parser.GetUserId()

	resp := a.service.GetItemList(req)
	if !resp.Success {
		return Fail(resp.Message, resp.Error.Error(), ctx)
	}

	return OkWithMessage(resp.Message, resp.Data, ctx)
}

// @summary 처방전 의약품 상세 조회
// @description 처방전 상세정보를 조회하는 엔드포인트
// @produce json
// @param id path string true "처방전 의약품 고유번호"
// @router /prescription/items/:id [get]
// @Security Bearer
func (a *prescriptionApi) GetItemById(ctx *fiber.Ctx) error {
	req := new(prescription.GetItemByIdRequest)
	parser := ParseRequest(req, PARAM, a.jwtClient, ctx)
	if parser.Error() != nil {
		return parser.Error()
	}

	req.UserId = parser.GetUserId()
	resp := a.service.GetItemById(req)
	if !resp.Success {
		return Fail(resp.Message, resp.Error.Error(), ctx)
	}

	return OkWithMessage(resp.Message, resp.Data, ctx)
}

// @summary 처방전 의약품 등록
// @description 처방전 의약품을 등록하는 엔드포인트
// @produce json
// @param dto body prescription.AddItemRequest true "처방전 의약품 등록정보"
// @router /prescription/items [post]
// @Security Bearer
func (a *prescriptionApi) AddItem(ctx *fiber.Ctx) error {
	req := new(prescription.AddItemRequest)
	parser := ParseRequest(req, BODY, a.jwtClient, ctx)
	if parser.Error() != nil {
		return parser.Error()
	}

	resp := a.service.AddItem(req)
	if !resp.Success {
		return response.New(response.CODE_FAIL_ADDITEM_PRESCRIPTION).
			SetMessage(resp.Message).
			SetErrors(resp.Error).
			Error(ctx)
	}

	return OkWithMessage(resp.Message, nil, ctx)
}

// @summary 처방전 의약품 업데이트
// @description 처방전 의약품을 업데이트하는 엔드포인트
// @produce json
// @param dto body prescription.UpdateRequest true "처방전 의약품 정보"
// @router /prescription/items [patch]
// @Security Bearer
func (a *prescriptionApi) UpdateItem(ctx *fiber.Ctx) error {
	req := new(prescription.UpdateItemRequest)
	parser := ParseRequest(req, BODY, a.jwtClient, ctx)
	if parser.Error() != nil {
		return parser.Error()
	}

	resp := a.service.UpdateItem(req)
	if !resp.Success {
		return response.New(response.CODE_FAIL_UPDATEITEM_PRESCRIPTION).
			SetMessage(resp.Message).
			SetErrors(resp.Error).
			Error(ctx)
	}

	return OkWithMessage(resp.Message, nil, ctx)
}

// @summary 처방전 의약품 삭제
// @description 처방전 의약품을 삭제하는 엔드포인트
// @produce json
// @param id path string true "처방전 의약품 고유번호"
// @router /prescription/items/:id [delete]
// @Security Bearer
func (a *prescriptionApi) DeleteItem(ctx *fiber.Ctx) error {
	req := new(prescription.DeleteItemRequest)
	parser := ParseRequest(req, BODY, a.jwtClient, ctx)
	if parser.Error() != nil {
		return parser.Error()
	}

	resp := a.service.DeleteItem(req)
	if !resp.Success {
		return response.New(response.CODE_FAIL_DELETEITEM_PRESCRIPTION).
			SetMessage(resp.Message).
			SetErrors(resp.Error).
			Error(ctx)
	}

	return OkWithMessage(resp.Message, nil, ctx)
}
