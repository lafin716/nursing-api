package plan

import (
	"github.com/google/uuid"
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
	plan             Repository
	prescriptionRepo prescription.Repository
	timezoneRepo     timezone.Repository
	timezoneLinkRepo timezonelink.Repository
	takeHistoryRepo  takehistory.Repository
	medicineRepo     medicine.Repository
}

func NewService(
	db *database.DatabaseClient,
	mono *mono.Client,
	plan Repository,
	prescriptionRepo prescription.Repository,
	timezoneRepo timezone.Repository,
	timezoneLinkRepo timezonelink.Repository,
	takeHistoryRepo takehistory.Repository,
	medicineRepo medicine.Repository,
) UseCase {
	return &planService{
		db:               db,
		mono:             mono,
		plan:             plan,
		prescriptionRepo: prescriptionRepo,
		timezoneRepo:     timezoneRepo,
		timezoneLinkRepo: timezoneLinkRepo,
		takeHistoryRepo:  takeHistoryRepo,
		medicineRepo:     medicineRepo,
	}
}

func (p *planService) Add(req *AddPlanRequest) *AddPlanResponse {
	txErr := p.db.BeginTx()
	if txErr != nil {
		return FailAddPlan("트랜잭션 생성 실패", txErr)
	}

	// 처방전 등록
	startedAt, err := p.mono.Date.Parse("Y-m-d", req.StartedAt)
	if txErr != nil {
		return FailAddPlan("날짜형식이 유효하지 않습니다", err)
	}
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

func (p *planService) Delete(req *DeletePlanRequest) *DeletePlanResponse {
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

func (p *planService) Take(req *TakeToggleRequest) *TakeToggleResponse {
	p.db.BeginTx()
	dateTime, err := p.mono.Date.Parse("Y-m-d H:i:s", req.TargetDate)
	if err != nil {
		return FailTakeToggle("날짜형식이 유효하지 않습니다", err)
	}

	// 복용여부
	var isTaken bool

	// 날짜기준 유효한 처방전 조회
	psList, err := p.prescriptionRepo.GetListBySearch(&prescription.SearchCondition{
		UserId:     req.UserId,
		TargetDate: dateTime,
	})
	if err != nil {
		return FailTakeToggle("해당 날짜에 등록된 처방전이 없습니다", err)
	}

	// 복용내역 조회
	truncatedDate := p.mono.Date.TruncateToDate(dateTime)
	tz, err := p.takeHistoryRepo.GetByTimezoneId(req.UserId, req.TimezoneId, truncatedDate)

	// 있는 경우 복용 취소 처리
	if tz != nil || err == nil {
		isTaken = false
		p.takeHistoryRepo.Delete(tz.ID)
		p.takeHistoryRepo.DeleteItemByHistoryId(tz.ID)
		p.db.CommitAndEnd()
		return OkTakeToggle(isTaken)
	}

	// 없는 경우 복용처리
	isTaken = true
	newHistory := &takehistory.TakeHistory{
		UserId:     req.UserId,
		TimezoneId: req.TimezoneId,
		TakeDate:   dateTime,
		TakeStatus: takehistory.DONE,
		CreatedAt:  time.Now(),
	}
	saved, err := p.takeHistoryRepo.Add(newHistory)
	if saved == nil || err != nil {
		p.db.RollbackAndEnd()
		return FailTakeToggle("복용처리 중 오류가 발생하였습니다", err)
	}

	// 처방전 고유번호 슬라이스 가공
	var prescriptionIds []uuid.UUID
	for _, ps := range psList {
		prescriptionIds = append(prescriptionIds, ps.ID)
	}

	// 타임존링크 조회
	tzlList, err := p.timezoneLinkRepo.GetByTimezoneIdAndPrescriptionIds(req.TimezoneId, prescriptionIds)
	if err != nil {
		p.db.RollbackAndEnd()
		return FailTakeToggle("복용시간대 조회 중 오류가 발생하였습니다", err)
	}

	var tzlSerial []uuid.UUID
	for _, tzl := range tzlList {
		tzlSerial = append(tzlSerial, tzl.ID)
	}

	// 복용아이템 조회
	items, err := p.prescriptionRepo.GetItemListByTimezoneLinkIds(tzlSerial)
	if err != nil {
		p.db.RollbackAndEnd()
		return FailTakeToggle("복용아이템 조회 중 오류가 발생하였습니다", err)
	}

	// 복용아이템 이력 추가
	for _, item := range items {
		newItem := &takehistory.TakeHistoryItem{
			UserId:             req.UserId,
			TakeHistoryId:      saved.ID,
			PrescriptionItemId: item.ID,
			TakeStatus:         takehistory.Y,
			TakeAmount:         item.TakeAmount,
			TakeUnit:           item.MedicineUnit,
			TakeDate:           dateTime,
			CreatedAt:          item.CreatedAt,
		}
		result, err := p.takeHistoryRepo.AddItem(newItem)
		if !result || err != nil {
			p.db.RollbackAndEnd()
			return FailTakeToggle("복용아이템 기록 중 오류가 발생하였습니다", err)
		}
	}

	p.db.CommitAndEnd()
	return OkTakeToggle(isTaken)
}

func (p *planService) PillToggle(req *PillToggleRequest) *PillToggleResponse {
	item, err := p.takeHistoryRepo.GetItemById(req.UserId, req.PillId)
	if err != nil {
		return FailPillToggle("복용 의약품을 찾을 수 없습니다", err)
	}

	isTaken := takehistory.Y == item.TakeStatus
	if isTaken {
		item.TakeStatus = takehistory.N
	} else {
		item.TakeStatus = takehistory.Y
	}
	saved, err := p.takeHistoryRepo.UpdateItem(item)
	if !saved || err != nil {
		return FailPillToggle("의약품 복용 처리 중 오류가 발생하였습니다", err)
	}

	return OkPillToggle(!isTaken)
}

func (p *planService) GetByMonth(req *GetByMonthRequest) *GetByMonthResponse {
	panic("")
}

func (p *planService) GetByDate(req *GetByDateRequest) *GetByDateResponse {
	// 날짜 세팅
	currentDate, err := p.mono.Date.Parse("Y-m-d", req.CurrentDate)
	if err != nil {
		return FailGetByDate("날짜형식이 유효하지 않습니다", err)
	}

	plans, err := p.plan.GetPlans(req.UserId, currentDate)
	if err != nil {
		return FailGetByDate("복용계획 조회 중 오류가 발생하였습니다.", err)
	}

	// 당일 복용내역 조회
	if err != nil {
		return FailGetByDate("복용내역 조회 중 오류가 발생하였습니다.", err)
	}

	// 당일 복용 계획 생성
	takePlan := &TakePlan{
		CurrentDate: p.mono.Date.Format(currentDate, "Y-m-d"),
		Plans:       plans,
	}
	return OkGetByDate(takePlan)
}

func (p *planService) UpdateMemo(req *UpdateMemoRequest) *UpdateMemoResponse {
	dateTime, err := p.mono.Date.Parse("Y-m-d", req.Date)
	if err != nil {
		return FailUpdateMemo("날짜 형식이 아닙니다", err)
	}
	tz, err := p.takeHistoryRepo.GetByTimezoneId(req.UserId, req.TimezoneId, dateTime)
	if err != nil {
		return FailUpdateMemo("복용처리 이후 메모를 업데이트 할 수 있습니다", err)
	}

	tz.Memo = req.Memo
	result, err := p.takeHistoryRepo.Update(tz)
	if !result || err != nil {
		return FailUpdateMemo("메모 업데이트 중 오류가 발생하였습니다", err)
	}

	return OkUpdateMemo()
}
