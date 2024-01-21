package takehistory

import (
	"nursing_api/internal/domain/prescription"
	"time"
)

type takeHistoryService struct {
	repo             TakeHistoryRepository
	prescriptionRepo prescription.PrescriptionRepository
}

func NewTakeHistoryService(
	repo TakeHistoryRepository,
	prescriptionRepo prescription.PrescriptionRepository,
) TakeHistoryUseCase {
	return &takeHistoryService{
		repo:             repo,
		prescriptionRepo: prescriptionRepo,
	}
}

func (t takeHistoryService) GetList(req *GetListRequest) *GetListResponse {
	list, err := t.repo.GetList(req.UserId)
	if err != nil {
		return FailGetList("복용내역 목록 조회 중 오류가 발생하였습니다.", err)
	}

	return OkGetList(list)
}

func (t takeHistoryService) GetDetail(req *GetDetailRequest) *GetDetailResponse {
	//TODO implement me
	panic("implement me")
}

func (t takeHistoryService) TakePlanToggle(req *TakePlanRequest) *TakePlanResponse {
	found, err := t.repo.GetByToday(req.UserId, req.TakeDate)
	if err != nil {
		return FailTakePillPlan("복용처리가 실패하였습니다.", err)
	}

	// 복용 완료상태인 경우
	if found.TakeStatus == TS_DONE {
		found.TakeStatus = ""
	} else {
		found.TakeStatus = TS_DONE
		found.TakeDate = time.Now()
	}
	found.UpdatedAt = time.Now()
	result, err := t.repo.Update(found)
	if err != nil {
		return FailTakePillPlan("복용처리 중 오류가 발생하였습니다.", err)
	}

	if !result {
		return OkTakePillPlan("복용처리가 실패하였습니다.")
	}

	return OkTakePillPlan("복용처리가 완료되었습니다.")
}

func (t takeHistoryService) TakePillToggle(req *TakePillRequest) *TakePillResponse {
	//TODO implement me
	panic("implement me")
}
