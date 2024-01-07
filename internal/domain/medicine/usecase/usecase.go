package usecase

import dto "nursing_api/internal/domain/medicine/dto"

type MedicineUseCase interface {
	SearchMedicine(pillName string) *dto.SearchPillResponse
}
