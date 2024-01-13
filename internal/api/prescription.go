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

func (a *prescriptionApi) GetList(ctx *fiber.Ctx) error {
	userId, err := getUserId(a.jwtClient, ctx)
	if err != nil {
		return err
	}

	resp := a.service.GetList(&prescription.GetListRequest{UserId: userId})
	if !resp.Success {
		return response.New(response.CODE_NO_DATA).SetErrors(err).Error(ctx)
	}

	return response.New(response.CODE_SUCCESS).
		SetData(resp.Data).
		Ok(ctx)
}

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
		return response.New(response.CODE_INVALID_PARAM).
			SetMessage(resp.Message).
			SetErrors(resp.Error).
			Error(ctx)
	}

	return response.New(response.CODE_SUCCESS).
		SetMessage(resp.Message).
		Ok(ctx)
}
