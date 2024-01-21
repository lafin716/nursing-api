package plan

import (
	"nursing_api/internal/domain/prescription"
	takehistory "nursing_api/internal/domain/take_history"
)

type planService struct {
	prescriptionRepo prescription.PrescriptionRepository
	takeHistoryRepo  takehistory.TakeHistoryRepository
}

func (p planService) GetByMonth(req *GetByMonthRequest) *GetByMonthResponse {
	//TODO implement me
	panic("implement me")
}

func (p planService) GetByDate(req *GetByDateRequest) *GetByDateResponse {
	//TODO implement me
	panic("implement me")
}

func (p planService) TakePlan(req *TakePlanRequest) *TakePlanResponse {
	//TODO implement me
	panic("implement me")
}

func (p planService) UnTakePlan(req *UnTakePlanRequest) *UnTakePlanResponse {
	//TODO implement me
	panic("implement me")
}

func (p planService) TakePill(req *TakePillRequest) *TakePillResponse {
	//TODO implement me
	panic("implement me")
}

func (p planService) UnTakePill(req *UnTakePillRequest) *UnTakePillResponse {
	//TODO implement me
	panic("implement me")
}

func (p planService) UpdateMemo(req *UpdateMemoRequest) *UpdateMemoResponse {
	//TODO implement me
	panic("implement me")
}

func NewPlanService(
	prescriptionRepo prescription.PrescriptionRepository,
	takeHistoryRepo takehistory.TakeHistoryRepository,
) PlanUseCase {
	return &planService{
		prescriptionRepo: prescriptionRepo,
		takeHistoryRepo:  takeHistoryRepo,
	}
}
