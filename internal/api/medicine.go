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

// TODO 검색어 2자 이상 유효성검사 구조체로 변경
func (u *medicineHttpApi) Search(ctx *fiber.Ctx) error {
	pillName := ctx.Params("pillName", "")
	if strings.TrimSpace(pillName) == "" {
		return response.New(response.CODE_INVALID_PARAM).BadRequest(ctx)
	}

	if len(pillName) < 2 {
		return response.New(response.CODE_INVALID_PARAM).
			SetMessage("검색어는 2자 이상 입력하세요.").
			BadRequest(ctx)
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
