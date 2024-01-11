package medicine

import medicine_api "nursing_api/pkg/api/medicine"

type medicineService struct {
	medicineApi medicine_api.MedicineApi
}

func (m medicineService) SearchMedicine(pillName string) *SearchPillResponse {
	resp := m.medicineApi.Summary().RetrievePill(pillName)
	if resp == nil {
		return FailSearchPill("의약품 정보를 찾을 수 없습니다.", nil)
	}

	medicines := []Medicine{}
	for _, item := range resp.Body.Items {
		medicines = append(medicines, Medicine{
			Name:       item.ItemName,
			EffectMemo: item.AtpnQesitm,
		})
	}

	return OkSearchPill(medicines)
}

func NewMedicineService(medicineApi medicine_api.MedicineApi) MedicineUseCase {
	return &medicineService{
		medicineApi: medicineApi,
	}
}
