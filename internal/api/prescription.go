package api

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/common/response"
	"nursing_api/internal/domain/prescription"
)

type PrescriptionApi interface {
	Regist(ctx *fiber.Ctx) error
}

type prescriptionApi struct {
	service prescription.PrescriptionUseCase
}

func NewPrescriptionApi(
	service prescription.PrescriptionUseCase,
) PrescriptionApi {
	return &prescriptionApi{
		service: service,
	}
}

func (a *prescriptionApi) Regist(ctx *fiber.Ctx) error {

	return response.New(response.CODE_SUCCESS).
		Ok(ctx)
}
