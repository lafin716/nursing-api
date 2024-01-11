package api

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/common/response"
	"nursing_api/internal/domain/medicine"
	"strings"
)

type MedicineHttpApi interface {
	Search(ctx *fiber.Ctx) error
}

type medicineHttpApi struct {
	usecase medicine.MedicineUseCase
}

func NewMedicineHttpApi(
	usecase medicine.MedicineUseCase,
) MedicineHttpApi {
	return &medicineHttpApi{
		usecase: usecase,
	}
}

func (u *medicineHttpApi) Search(ctx *fiber.Ctx) error {
	pillName := ctx.Params("pillName", "")
	if strings.TrimSpace(pillName) == "" {
		return response.New(response.CODE_INVALID_PARAM).BadRequest(ctx)
	}

	result := u.usecase.SearchMedicine(pillName)
	if !result.Success {
		return response.New(response.CODE_ERROR).
			SetMessage(result.Message).
			SetErrors(result.Error).
			Error(ctx)
	}

	return response.New(response.CODE_SUCCESS).
		SetData(result.Pills).
		Ok(ctx)
}
