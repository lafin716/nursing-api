package plan

import (
	"log"
	"nursing_api/internal/domain/prescription"
	takehistory "nursing_api/internal/domain/take_history"
	"time"
)

type UseCase interface {
	Add(req *AddPlanRequest) *AddPlanResponse
	GetByMonth(req *GetByMonthRequest) *GetByMonthResponse
	GetByDate(req *GetByDateRequest) *GetByDateResponse
	Take(req *TakeToggleRequest) *TakeToggleResponse
	PillToggle(req *PillToggleRequest) *PillToggleResponse
	UpdateMemo(req *UpdateMemoRequest) *UpdateMemoResponse
}

const (
	DATE_LAYOUT = "2006-01-02"
)

type planService struct {
	prescriptionRepo prescription.Repository
	takeHistoryRepo  takehistory.Repository
}

func NewService(
	prescriptionRepo prescription.Repository,
	takeHistoryRepo takehistory.Repository,
) UseCase {
	return &planService{
		prescriptionRepo: prescriptionRepo,
		takeHistoryRepo:  takeHistoryRepo,
	}
}

func (p planService) Add(req *AddPlanRequest) *AddPlanResponse {
	//TODO implement me
	panic("implement me")
}

func (p planService) Take(req *TakeToggleRequest) *TakeToggleResponse {
	//TODO implement me
	panic("implement me")
}

func (p planService) PillToggle(req *PillToggleRequest) *PillToggleResponse {
	//TODO implement me
	panic("implement me")
}

func (p planService) GetByMonth(req *GetByMonthRequest) *GetByMonthResponse {
	panic("")
}

func (p planService) GetByDate(req *GetByDateRequest) *GetByDateResponse {
	// 날짜 세팅
	currentDate, err := time.Parse(DATE_LAYOUT, req.CurrentDate)
	if err != nil {
		currentDate, _ = time.Parse(DATE_LAYOUT, time.Now().Format(DATE_LAYOUT))
	}

	// 당일 처방전 조회
	search := &prescription.SearchCondition{
		UserId:     req.UserId,
		TargetDate: currentDate,
		Limit:      10,
	}

	log.Printf("처방전 파라미터: %+v", search)
	origin, err := p.prescriptionRepo.GetItemListBySearch(search)
	if err != nil {
		return FailGetByDate("처방전 조회 중 오류가 발생하였습니다.", err)
	}

	for _, item := range origin {
		log.Printf("처방전 목록: %+v", item)
	}

	// 당일 복용내역 조회
	if err != nil {
		return FailGetByDate("복용내역 조회 중 오류가 발생하였습니다.", err)
	}

	// 당일 복용 계획 생성
	plan := &TakePlan{
		CurrentDate: currentDate.Format(DATE_LAYOUT),
	}
	return OkGetByDate(plan)
}

func (p planService) UpdateMemo(req *UpdateMemoRequest) *UpdateMemoResponse {
	//TODO implement me
	panic("implement me")
}
