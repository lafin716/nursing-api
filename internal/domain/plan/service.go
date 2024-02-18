package plan

import (
	"github.com/google/uuid"
	"log"
	"nursing_api/internal/domain/medicine"
	"nursing_api/internal/domain/prescription"
	takehistory "nursing_api/internal/domain/take_history"
	"nursing_api/internal/domain/timezone"
	timezonelink "nursing_api/internal/domain/timezone_link"
	"nursing_api/pkg/database"
	"nursing_api/pkg/mono"
	"time"
)

type UseCase interface {
	Add(req *AddPlanRequest) *AddPlanResponse
	Delete(req *DeletePlanRequest) *DeletePlanResponse
	GetByMonth(req *GetByMonthRequest) *GetByMonthResponse
	GetByDate(req *GetByDateRequest) *GetByDateResponse
	Take(req *TakeToggleRequest) *TakeToggleResponse
	PillToggle(req *PillToggleRequest) *PillToggleResponse
	UpdateMemo(req *UpdateMemoRequest) *UpdateMemoResponse
}

type planService struct {
	db               *database.DatabaseClient
	mono             *mono.Client
	prescriptionRepo prescription.Repository
	timezoneRepo     timezone.Repository
	timezoneLinkRepo timezonelink.Repository
	takeHistoryRepo  takehistory.Repository
	medicineRepo     medicine.Repository
}

func NewService(
	db *database.DatabaseClient,
	mono *mono.Client,
	prescriptionRepo prescription.Repository,
	timezoneRepo timezone.Repository,
	timezoneLinkRepo timezonelink.Repository,
	takeHistoryRepo takehistory.Repository,
	medicineRepo medicine.Repository,
) UseCase {
	return &planService{
		db:               db,
		mono:             mono,
		prescriptionRepo: prescriptionRepo,
		timezoneRepo:     timezoneRepo,
		timezoneLinkRepo: timezoneLinkRepo,
		takeHistoryRepo:  takeHistoryRepo,
		medicineRepo:     medicineRepo,
	}
}

func (p planService) Add(req *AddPlanRequest) *AddPlanResponse {
	txErr := p.db.BeginTx()
	if txErr != nil {
		return FailAddPlan("트랜잭션 생성 실패", txErr)
	}

	// 처방전 등록
	startedAt := p.mono.Date.ParseForce("Y-m-d", req.StartedAt)
	finishedAt := startedAt.AddDate(0, 0, req.TakeDays)
	psReq := &prescription.Prescription{
		UserId:           req.UserId,
		PrescriptionName: req.Name,
		HospitalName:     req.Hospital,
		TakeDays:         req.TakeDays,
		StartedAt:        startedAt,
		FinishedAt:       finishedAt,
		Memo:             req.Memo,
	}
	savedPs, err := p.prescriptionRepo.Add(psReq)
	if err != nil {
		return FailAddPlan("처방전 등록 실패", err)
	}

	// 타임존 링크 연결
	for _, tz := range req.Timezones {
		// 타임존 조회
		foundTz, err := p.timezoneRepo.GetTimeZone(tz.TimezoneId, req.UserId)
		if err != nil {
			p.db.RollbackAndEnd()
			return FailAddPlan("복용시간대를 찾을 수 없습니다", err)
		}

		// 현재는 기본 타임존의 정보를 DB에서 읽어서 저장
		tlReq := &timezonelink.TimeZoneLink{
			PrescriptionId: savedPs.ID,
			TimeZoneId:     foundTz.ID,
			TimeZoneName:   foundTz.Name,
			UseAlert:       false,
			Midday:         foundTz.Midday,
			Hour:           foundTz.Hour,
			Minute:         foundTz.Minute,
			CreatedAt:      time.Now(),
		}
		savedTimezoneLink, err := p.timezoneLinkRepo.ConnectPrescription(tlReq)
		if err != nil {
			p.db.RollbackAndEnd()
			return FailAddPlan("복용시간대 연결이 실패하였습니다", err)
		}

		// 의약품 등록
		for _, mdc := range tz.Medicines {
			// 현재는 기본 의약품의 정보를 DB에서 읽어서 저장
			foundMdc, err := p.medicineRepo.GetById(mdc.MedicineId)
			if err != nil {
				p.db.RollbackAndEnd()
				return FailAddPlan("의약품을 찾을 수 없습니다", err)
			}

			psiReq := &prescription.PrescriptionItem{
				TimeZoneLinkId: savedTimezoneLink.ID,
				MedicineId:     mdc.MedicineId,
				MedicineName:   foundMdc.Name,
				TakeAmount:     mdc.TakeAmount,
				MedicineUnit:   mdc.TakeUnit,
				Memo:           mdc.Memo,
				CreatedAt:      time.Now(),
			}
			_, err = p.prescriptionRepo.AddItem(psiReq)
			if err != nil {
				p.db.RollbackAndEnd()
				return FailAddPlan("복용시간대 아이템 추가가 실패하였습니다", err)
			}
		}
	}

	p.db.CommitAndEnd()
	return OkAddPlan()
}

func (p planService) Delete(req *DeletePlanRequest) *DeletePlanResponse {
	p.db.BeginTx()

	// 복용계획 조회
	ps, err := p.prescriptionRepo.GetById(req.ID)
	if err != nil {
		return FailDeletePlan("복용계획을 찾을 수 없습니다", err)
	}

	// 복용계획 삭제처리
	result, err := p.prescriptionRepo.Delete(ps.ID)
	if !result || err != nil {
		p.db.RollbackAndEnd()
		return FailDeletePlan("복용계획 삭제가 실패하였습니다", err)
	}

	// 복용시간대 링크 조회
	tzls, err := p.timezoneLinkRepo.GetByPrescription(ps.ID)
	// 없을 수도 있으므로 없는경우 정상 응답
	if err == nil {
		p.db.CommitAndEnd()
		return OkDeletePlan()
	}

	// 시간대 고유번호를 배열화
	tzlIds := []uuid.UUID{}
	for _, tzl := range tzls {
		tzlIds = append(tzlIds, tzl.ID)
	}

	// 시간대 삭제처리
	result, err = p.timezoneLinkRepo.DeleteByIds(tzlIds)
	if !result || err != nil {
		p.db.RollbackAndEnd()
		return FailDeletePlan("복용 타임존 삭제 중 오류가 발생하였습니다", err)
	}

	// 복용계획 아이템 조회
	psItems, err := p.prescriptionRepo.GetItemListByTimezoneLinkIds(tzlIds)
	// 없을 수도 있으므로 없는 경우 정상응답
	if err != nil {
		p.db.CommitAndEnd()
		return OkDeletePlan()
	}

	// 복용계획 아이템 삭제처리
	psItemIds := []uuid.UUID{}
	for _, psItem := range psItems {
		psItemIds = append(psItemIds, psItem.ID)
	}
	result, err = p.prescriptionRepo.DeleteItemByIds(psItemIds)
	if !result || err != nil {
		p.db.RollbackAndEnd()
		return FailDeletePlan("복용 의약품 삭제 중 오류가 발생하였습니다", err)
	}

	p.db.CommitAndEnd()
	return OkDeletePlan()
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
	currentDate, err := p.mono.Date.Parse("Y-m-d", req.CurrentDate)
	if err != nil {
		currentDate = time.Now()
	}

	// 당일 처방전 조회
	search := &prescription.SearchCondition{
		UserId:     req.UserId,
		TargetDate: currentDate,
		Limit:      10,
	}

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
		CurrentDate: p.mono.Date.Format(currentDate, "Y-m-d"),
	}
	return OkGetByDate(plan)
}

func (p planService) UpdateMemo(req *UpdateMemoRequest) *UpdateMemoResponse {
	//TODO implement me
	panic("implement me")
}
