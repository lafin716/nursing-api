package service

import (
	domain "nursing_api/internal/domain/medicine"
	dto "nursing_api/internal/domain/medicine/dto"
	"nursing_api/internal/domain/medicine/usecase"
	"nursing_api/pkg/api/medicine"
)

type medicineService struct {
	medicineApi medicine.MedicineApi
}

func (m medicineService) SearchMedicine(pillName string) *dto.SearchPillResponse {
	resp := m.medicineApi.Summary().RetrievePill(pillName)
	if resp == nil {
		return dto.FailSearchPill("의약품 정보를 찾을 수 없습니다.", nil)
	}

	medicines := []domain.Medicine{}
	for _, item := range resp.Body.Items {
		medicines = append(medicines, domain.Medicine{
			Name:       item.ItemName,
			EffectMemo: item.AtpnQesitm,
		})
	}

	return dto.OkSearchPill(medicines)
}

func NewMedicineService(medicineApi medicine.MedicineApi) usecase.MedicineUseCase {
	return &medicineService{
		medicineApi: medicineApi,
	}
}
