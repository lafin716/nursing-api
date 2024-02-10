package api

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/common/response"
	"nursing_api/internal/domain/prescription"
	"nursing_api/pkg/jwt"
)

type PrescriptionApi interface {
	Regist(ctx *fiber.Ctx) error
	GetList(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	AddItem(ctx *fiber.Ctx) error
	UpdateItem(ctx *fiber.Ctx) error
	DeleteItem(ctx *fiber.Ctx) error
}

type prescriptionApi struct {
	service   prescription.PrescriptionUseCase
	jwtClient *jwt.JwtClient
}

func NewPrescriptionApi(
	service prescription.PrescriptionUseCase,
	jwtClient *jwt.JwtClient,
) PrescriptionApi {
	return &prescriptionApi{
		service:   service,
		jwtClient: jwtClient,
	}
}

// 처방전 리스트
func (a *prescriptionApi) GetList(ctx *fiber.Ctx) error {
	req := new(prescription.GetListRequest)
	err := ctx.QueryParser(req)
	if err != nil {
		return response.New(response.CODE_INVALID_PARAM).SetErrors(err).Error(ctx)
	}

	userId, err := getUserId(a.jwtClient, ctx)
	if err != nil {
		return err
	}
	req.UserId = userId

	resp := a.service.GetList(req)
	if !resp.Success {
		return response.New(response.CODE_NO_DATA).SetErrors(err).Error(ctx)
	}

	return response.New(response.CODE_SUCCESS).
		SetData(resp.Data).
		SetMessage(resp.Message).
		Ok(ctx)
}

// 처방전 등록
func (a *prescriptionApi) Regist(ctx *fiber.Ctx) error {
	userId, err := getUserId(a.jwtClient, ctx)
	if err != nil {
		return err
	}

	req := new(prescription.RegisterRequest)
	err = ctx.BodyParser(req)
	if err != nil {
		return response.New(response.CODE_INVALID_PARAM).
			SetErrors(err).
			Error(ctx)
	}

	req.UserId = userId
	errs := validateParameter(req)
	if errs != nil {
		return response.New(response.CODE_INVALID_PARAM).
			SetErrors(errs).
			Error(ctx)
	}

	resp := a.service.Register(req)
	if !resp.Success {
		return response.New(response.CODE_FAIL_ADD_PRESCRIPTION).
			SetMessage(resp.Message).
			SetErrors(resp.Error).
			Error(ctx)
	}

	return response.New(response.CODE_SUCCESS).
		SetMessage(resp.Message).
		Ok(ctx)
}

// 처방전 업데이트
func (a *prescriptionApi) Update(ctx *fiber.Ctx) error {
	req := new(prescription.UpdateRequest)
	err := ctx.BodyParser(req)
	if err != nil {
		return response.New(response.CODE_INVALID_PARAM).
			SetErrors(err.Error()).
			Error(ctx)
	}

	errs := validateParameter(req)
	if errs != nil {
		return response.New(response.CODE_INVALID_PARAM).
			SetErrors(errs).
			Error(ctx)
	}

	resp := a.service.Update(req)
	if !resp.Success {
		return response.New(response.CODE_FAIL_UPDATE_PRESCRIPTION).
			SetMessage(resp.Message).
			SetErrors(resp.Error).
			Error(ctx)
	}

	return response.New(response.CODE_SUCCESS).
		SetMessage(resp.Message).
		SetData(resp.Data).
		Ok(ctx)
}

// 처방전 삭제
func (a *prescriptionApi) Delete(ctx *fiber.Ctx) error {
	req := new(prescription.DeleteRequest)
	err := ctx.ParamsParser(req)
	if err != nil {
		return response.New(response.CODE_INVALID_PARAM).
			SetErrors(err.Error()).
			Error(ctx)
	}

	resp := a.service.Delete(req)
	if !resp.Success {
		return response.New(response.CODE_FAIL_DELETE_PRESCRIPTION).
			SetMessage(resp.Message).
			SetErrors(resp.Error).
			Error(ctx)
	}

	return response.New(response.CODE_SUCCESS).
		SetMessage(resp.Message).
		Ok(ctx)
}

// 처방전 의약품 추가
func (a *prescriptionApi) AddItem(ctx *fiber.Ctx) error {
	req := new(prescription.AddItemRequest)
	err := ctx.BodyParser(req)
	if err != nil {
		return FailParam(err.Error(), ctx)
	}

	errs := validateParameter(req)
	if errs != nil {
		return FailParam(errs, ctx)
	}

	resp := a.service.AddItem(req)
	if !resp.Success {
		return response.New(response.CODE_FAIL_ADDITEM_PRESCRIPTION).
			SetMessage(resp.Message).
			SetErrors(resp.Error).
			Error(ctx)
	}

	return response.New(response.CODE_SUCCESS).
		SetMessage(resp.Message).
		Ok(ctx)
}

// 처방전 의약품 업데이트
func (a *prescriptionApi) UpdateItem(ctx *fiber.Ctx) error {
	req := new(prescription.UpdateItemRequest)
	err := ctx.BodyParser(req)
	if err != nil {
		return response.New(response.CODE_INVALID_PARAM).
			SetErrors(err.Error()).
			Error(ctx)
	}

	errs := validateParameter(req)
	if errs != nil {
		return response.New(response.CODE_INVALID_PARAM).
			SetErrors(errs).
			Error(ctx)
	}

	resp := a.service.UpdateItem(req)
	if !resp.Success {
		return response.New(response.CODE_FAIL_UPDATEITEM_PRESCRIPTION).
			SetMessage(resp.Message).
			SetErrors(resp.Error).
			Error(ctx)
	}

	return response.New(response.CODE_SUCCESS).
		SetMessage(resp.Message).
		Ok(ctx)
}

// 처방전 아이템 삭제
func (a *prescriptionApi) DeleteItem(ctx *fiber.Ctx) error {
	req := new(prescription.DeleteItemRequest)
	err := ctx.ParamsParser(req)
	if err != nil {
		return response.New(response.CODE_INVALID_PARAM).
			SetErrors(err.Error()).
			Error(ctx)
	}

	errs := validateParameter(req)
	if errs != nil {
		return response.New(response.CODE_INVALID_PARAM).
			SetErrors(errs).
			Error(ctx)
	}

	resp := a.service.DeleteItem(req)
	if !resp.Success {
		return response.New(response.CODE_FAIL_DELETEITEM_PRESCRIPTION).
			SetMessage(resp.Message).
			SetErrors(resp.Error).
			Error(ctx)
	}

	return response.New(response.CODE_SUCCESS).
		SetMessage(resp.Message).
		Ok(ctx)
}
