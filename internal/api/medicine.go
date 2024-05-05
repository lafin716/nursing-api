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
	usecase medicine.UseCase
}

func NewMedicineHttpApi(
	usecase medicine.UseCase,
) MedicineHttpApi {
	return &medicineHttpApi{
		usecase: usecase,
	}
}

// @summary 의약품 검색
// @description 의약품을 검색하는 엔드포인트. 공공 API를 통해 조회하며, 1번 조회 시 DB에 캐싱처리함
// @param pillName path string true "의약품 이름(like 검색)"
// @router /medicine/search/{pillName} [get]
// @Security Bearer
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

	result := u.usecase.Search(pillName)
	return ResolveResponse(result, ctx)
}
