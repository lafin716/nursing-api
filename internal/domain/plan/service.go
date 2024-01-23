package plan

import (
	"fmt"
	"log"
	"nursing_api/internal/domain/prescription"
	takehistory "nursing_api/internal/domain/take_history"
	"time"
)

const (
	DATE_LAYOUT = "2006-01-02"
)

type planService struct {
	prescriptionRepo prescription.PrescriptionRepository
	takeHistoryRepo  takehistory.TakeHistoryRepository
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
	search := &prescription.PrescriptionSearch{
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
	log, err := p.takeHistoryRepo.GetByToday(req.UserId, currentDate)
	if err != nil {
		return FailGetByDate("복용내역 조회 중 오류가 발생하였습니다.", err)
	}

	fmt.Printf("로그 : %+v", log)

	// 데이터 조합 및 가공처리
	var mapPills = map[string][]PillItem{}
	for _, pillItem := range origin {
		timeZone := pillItem.TakeTimeZone
		pill := PillItem{
			PillName:   pillItem.MedicineName,
			MedicineId: pillItem.MedicineId,
			TakeAmount: pillItem.TakeAmount,
		}

		plans := mapPills[timeZone]
		plans = append(plans, pill)
		mapPills[timeZone] = plans
	}

	var planItems []PlanItem
	for k, v := range mapPills {
		planItems = append(planItems, PlanItem{
			PlanName:     GetTimeZone(k),
			TakeTimeZone: k,
			Pills:        v,
		})
	}

	// 당일 복용 계획 생성
	plan := &TakePlan{
		CurrentDate: currentDate.Format(DATE_LAYOUT),
		Plans:       planItems,
	}
	return OkGetByDate(plan)
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
