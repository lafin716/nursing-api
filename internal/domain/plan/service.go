package plan

import (
	"github.com/google/uuid"
	"nursing_api/internal/common/dto"
	"nursing_api/internal/common/response"
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
	Add(req *AddPlanRequest) dto.BaseResponse[any]
	Delete(req *DeletePlanRequest) dto.BaseResponse[any]
	GetByMonth(req *GetByMonthRequest) dto.BaseResponse[any]
	GetByDate(req *GetByDateRequest) dto.BaseResponse[TakePlan]
	Take(req *TakeToggleRequest) dto.BaseResponse[bool]
	PillToggle(req *PillToggleRequest) dto.BaseResponse[bool]
	UpdateMemo(req *UpdateMemoRequest) dto.BaseResponse[any]
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

func (p *planService) Add(req *AddPlanRequest) dto.BaseResponse[any] {
	txErr := p.db.BeginTx()
	if txErr != nil {
		return dto.Fail[any](response.CODE_ERROR_TRANSACTION, txErr)
	}

	// 처방전 등록
	startedAt, err := p.mono.Date.Parse("Y-m-d", req.StartedAt)
	if txErr != nil {
		return dto.Fail[any](response.CODE_NOT_AVAILABLE_DATE, err)
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
		return dto.Fail[any](response.CODE_FAIL_ADD_PRESCRIPTION, err)
	}

	// 타임존 링크 연결
	for _, tz := range req.Timezones {
		// 타임존 조회
		foundTz, err := p.timezoneRepo.GetTimeZone(tz.TimezoneId, req.UserId)
		if err != nil {
			p.db.RollbackAndEnd()
			return dto.Fail[any](response.CODE_NOT_FOUND_PLAN_TIMEZONE, err)
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
			return dto.Fail[any](response.CODE_FAIL_CONNECT_TIMEZONE, err)
		}

		// 의약품 등록
		for _, mdc := range tz.Medicines {
			// 현재는 기본 의약품의 정보를 DB에서 읽어서 저장
			foundMdc, err := p.medicineRepo.GetById(mdc.MedicineId)
			if err != nil {
				p.db.RollbackAndEnd()
				return dto.Fail[any](response.CODE_NOT_FOUND_MEDICINE, err)
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
				return dto.Fail[any](response.CODE_FAIL_ADD_ITEM_PLAN, err)
			}
		}
	}

	p.db.CommitAndEnd()
	return dto.Ok[any](response.CODE_SUCCESS, nil)
}

func (p *planService) Delete(req *DeletePlanRequest) dto.BaseResponse[any] {
	p.db.BeginTx()

	// 복용계획 조회
	ps, err := p.prescriptionRepo.GetById(req.ID)
	if err != nil {
		return dto.Fail[any](response.CODE_NOT_FOUND_PLAN, err)
	}

	// 복용계획 삭제처리
	result, err := p.prescriptionRepo.Delete(ps.ID)
	if !result || err != nil {
		p.db.RollbackAndEnd()
		return dto.Fail[any](response.CODE_FAIL_DELETE_PLAN, err)
	}

	// 복용시간대 링크 조회
	tzls, err := p.timezoneLinkRepo.GetByPrescription(ps.ID)
	// 없을 수도 있으므로 없는경우 정상 응답
	if err == nil {
		p.db.CommitAndEnd()
		return dto.Ok[any](response.CODE_SUCCESS, nil)
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
		return dto.Fail[any](response.CODE_FAIL_DELETE_PLAN_TIMEZONE, nil)
	}

	// 복용계획 아이템 조회
	psItems, err := p.prescriptionRepo.GetItemListByTimezoneLinkIds(tzlIds)
	// 없을 수도 있으므로 없는 경우 정상응답
	if err != nil {
		p.db.CommitAndEnd()
		return dto.Ok[any](response.CODE_SUCCESS, nil)
	}

	// 복용계획 아이템 삭제처리
	psItemIds := []uuid.UUID{}
	for _, psItem := range psItems {
		psItemIds = append(psItemIds, psItem.ID)
	}
	result, err = p.prescriptionRepo.DeleteItemByIds(psItemIds)
	if !result || err != nil {
		p.db.RollbackAndEnd()
		return dto.Fail[any](response.CODE_FAIL_DELETE_ITEM_PLAN, err)
	}

	p.db.CommitAndEnd()
	return dto.Ok[any](response.CODE_SUCCESS, nil)
}

func (p *planService) Take(req *TakeToggleRequest) dto.BaseResponse[bool] {
	p.db.BeginTx()
	dateTime, err := p.mono.Date.Parse("Y-m-d H:i:s", req.TargetDate)
	if err != nil {
		return dto.Fail[bool](response.CODE_NOT_AVAILABLE_DATE, err)
	}

	// 복용여부
	var isTaken bool

	// 날짜기준 유효한 처방전 조회
	psList, err := p.prescriptionRepo.GetListBySearch(&prescription.SearchCondition{
		UserId:     req.UserId,
		TargetDate: dateTime,
	})
	if err != nil {
		return dto.Fail[bool](response.CODE_NO_RESULT_PLAN_BY_DATE, err)
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
		return dto.Ok[bool](response.CODE_SUCCESS, &isTaken)
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
		return dto.Fail[bool](response.CODE_FAIL_TAKE_PLAN, err)
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
		return dto.Fail[bool](response.CODE_NOT_FOUND_PLAN_TIMEZONE, err)
	}

	var tzlSerial []uuid.UUID
	for _, tzl := range tzlList {
		tzlSerial = append(tzlSerial, tzl.ID)
	}

	// 복용아이템 조회
	items, err := p.prescriptionRepo.GetItemListByTimezoneLinkIds(tzlSerial)
	if err != nil {
		p.db.RollbackAndEnd()
		return dto.Fail[bool](response.CODE_NOT_FOUND_PLAN_ITEM, err)
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
			return dto.Fail[bool](response.CODE_FAIL_SAVE_PLAN_LOG, err)
		}
	}

	p.db.CommitAndEnd()
	return dto.Ok[bool](response.CODE_SUCCESS, &isTaken)
}

func (p *planService) PillToggle(req *PillToggleRequest) dto.BaseResponse[bool] {
	item, err := p.takeHistoryRepo.GetItemById(req.UserId, req.PillId)
	if err != nil {
		return dto.Fail[bool](response.CODE_NOT_FOUND_PLAN_ITEM, err)
	}

	isTaken := takehistory.Y == item.TakeStatus
	if isTaken {
		item.TakeStatus = takehistory.N
	} else {
		item.TakeStatus = takehistory.Y
	}

	isTaken = !isTaken
	saved, err := p.takeHistoryRepo.UpdateItem(item)
	if !saved || err != nil {
		return dto.Fail[bool](response.CODE_FAIL_TAKE_PLAN, err)
	}

	return dto.Ok[bool](response.CODE_SUCCESS, &isTaken)
}

func (p *planService) GetByMonth(req *GetByMonthRequest) dto.BaseResponse[any] {
	// TODO 구현필요
	panic("")
}

func (p *planService) GetByDate(req *GetByDateRequest) dto.BaseResponse[TakePlan] {
	// 날짜 세팅
	currentDate, err := p.mono.Date.Parse("Y-m-d", req.CurrentDate)
	if err != nil {
		return dto.Fail[TakePlan](response.CODE_NOT_AVAILABLE_DATE, err)
	}

	// 해당 날짜에 등록된 처방 의약품 일괄 조회
	plans, err := p.plan.GetPlans(req.UserId, currentDate)
	if err != nil {
		return dto.Fail[TakePlan](response.CODE_NO_RESULT_PLAN_BY_DATE, err)
	}

	// 당일 복용 계획 생성
	takePlan := &TakePlan{
		CurrentDate: p.mono.Date.Format(currentDate, "Y-m-d"),
		Plans:       plans,
	}
	return dto.Ok[TakePlan](response.CODE_SUCCESS, takePlan)
}

func (p *planService) UpdateMemo(req *UpdateMemoRequest) dto.BaseResponse[any] {
	dateTime, err := p.mono.Date.Parse("Y-m-d", req.Date)
	if err != nil {
		return dto.Fail[any](response.CODE_NOT_AVAILABLE_DATE, err)
	}
	tz, err := p.takeHistoryRepo.GetByTimezoneId(req.UserId, req.TimezoneId, dateTime)
	if err != nil {
		return dto.Fail[any](response.CODE_CANNOT_UPDATE_MEMO_BEFORE_TAKING, err)
	}

	tz.Memo = req.Memo
	result, err := p.takeHistoryRepo.Update(tz)
	if !result || err != nil {
		return dto.Fail[any](response.CODE_FAIL_UPDATE_MEMO, err)
	}

	return dto.Ok[any](response.CODE_SUCCESS, nil)
}
