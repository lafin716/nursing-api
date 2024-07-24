package plan

import (
	"fmt"
	"nursing_api/internal/common/dto"
	"nursing_api/internal/common/response"
	"nursing_api/pkg/mono"
	"sort"
	"strconv"
	"time"
)

type UseCase interface {
	GetByDate(req *GetByDateRequest) dto.BaseResponse[TakePlan]
	GetByMonth(req *GetByMonthRequest) dto.BaseResponse[Summary]
	Add(req *AddPlanRequest) dto.BaseResponse[any]
	Delete(req *DeletePlanRequest) dto.BaseResponse[any]
	Take(req *TakeToggleRequest) dto.BaseResponse[bool]
	PillToggle(req *PillToggleRequest) dto.BaseResponse[bool]
	UpdateMemo(req *UpdateMemoRequest) dto.BaseResponse[any]
	PillTakeAmountUpdate(req *PillTakeAmountUpdateRequest) dto.BaseResponse[bool]
}

type planService struct {
	mono *mono.Client
	repo Repository
}

func NewService(
	mono *mono.Client,
	repo Repository,
) UseCase {
	return &planService{
		mono: mono,
		repo: repo,
	}
}

// 일별 복용계획 조회
func (p *planService) GetByDate(req *GetByDateRequest) dto.BaseResponse[TakePlan] {
	// 날짜 세팅
	currentDate, err := p.mono.Date.Parse("Y-m-d", req.CurrentDate)
	if err != nil {
		return dto.Fail[TakePlan](response.CODE_NOT_AVAILABLE_DATE, err)
	}

	// 시간대별 복용아이템 조회 및 그룹핑
	tzMap, err := p.repo.GetPlanMap(req.UserId, currentDate)
	if err != nil || tzMap == nil {
		return dto.Ok[TakePlan](response.CODE_SUCCESS, &TakePlan{
			CurrentDate: p.mono.Date.Format(currentDate, "Y-m-d"),
			Plans:       make([]Plan, 0),
		})
	}

	// 복용계획 정렬
	plans := make([]Plan, 0, len(tzMap))
	for _, v := range tzMap {
		plans = append(plans, *v)
	}
	sort.Slice(plans, func(i, j int) bool {
		org := (plans[i].Hour * 100) + plans[i].Minute
		dst := (plans[j].Hour * 100) + plans[j].Minute
		return org < dst
	})

	// 당일 복용 계획 생성
	takePlan := &TakePlan{
		CurrentDate: p.mono.Date.Format(currentDate, "Y-m-d"),
		Plans:       plans,
	}
	return dto.Ok[TakePlan](response.CODE_SUCCESS, takePlan)
}

// 월별 복용계획 조회
func (p *planService) GetByMonth(req *GetByMonthRequest) dto.BaseResponse[Summary] {
	yearNum, err := strconv.Atoi(req.Year)
	if err != nil {
		return dto.Fail[Summary](response.CODE_NOT_AVAILABLE_DATE, err)
	}
	monthNum, err := strconv.Atoi(req.Month)
	if err != nil {
		return dto.Fail[Summary](response.CODE_NOT_AVAILABLE_DATE, err)
	}

	// 월별 처방전 조회
	items, err := p.repo.GetPrescriptionInMonth(req.UserId, yearNum, monthNum)
	if err != nil {
		return dto.Fail[Summary](response.CODE_NO_RESULT_PLAN_BY_MONTH, err)
	}

	// 월별 복용계획 생성
	summary := &Summary{
		Year:  fmt.Sprintf("%d", yearNum),
		Month: fmt.Sprintf("%d", monthNum),
	}

	// 복용계획 아이템 추가
	for _, ps := range items {
		summaryItem := &SummaryItem{
			ID:           ps.ID,
			Name:         ps.PrescriptionName,
			HospitalName: ps.HospitalName,
			StartedAt:    p.mono.Date.Format(ps.StartedAt, "Y-m-d"),
			FinishedAt:   p.mono.Date.Format(ps.FinishedAt, "Y-m-d"),
			TakeDays:     ps.TakeDays,
		}
		summary.Items = append(summary.Items, *summaryItem)
	}

	return dto.Ok[Summary](response.CODE_SUCCESS, summary)
}

// 복용계획 추가
func (p *planService) Add(req *AddPlanRequest) dto.BaseResponse[any] {
	tx := p.repo.GetTxManager()
	err := tx.BeginTx()
	if err != nil {
		return dto.Fail[any](response.CODE_ERROR_TRANSACTION, err)
	}

	// 처방전 생성
	ps, err := p.repo.AddPrescription(*req)
	if err != nil {
		_ = tx.RollbackTx()
		return dto.Fail[any](response.CODE_FAIL_ADD_PRESCRIPTION, err)
	}

	startedAt, err := p.mono.Date.Parse("Y-m-d", req.StartedAt)
	if err != nil {
		_ = tx.RollbackTx()
		return dto.Fail[any](response.CODE_FAIL_ADD_TAKE_HISTORY_ITEM, err)
	}

	// 처방전 하위 타임존 순회
	for _, subTz := range req.Timezones {
		// 타임존 조회
		tz, err := p.repo.GetTimeZone(subTz.TimezoneId)
		if err != nil {
			_ = tx.RollbackTx()
			return dto.Fail[any](response.CODE_NOT_FOUND_TIMEZONE, err)
		}

		// 처방전 아이템 순회
		for _, subMedicine := range subTz.Medicines {
			// 처방전 아이템 총 수량 남은 수량 계산
			subMedicine.TotalAmount = float64(ps.TakeDays) * subMedicine.TakeAmount
			subMedicine.RemainAmount = float64(ps.TakeDays) * subMedicine.TakeAmount

			// 의약품 조회
			med, err := p.repo.GetMedicine(subMedicine.MedicineId)
			if err != nil {
				_ = tx.RollbackTx()
				return dto.Fail[any](response.CODE_NOT_FOUND_MEDICINE, err)
			}

			// 처방전 아이템 생성
			pscItem, err := p.repo.AddPrescriptionItems(ps, tz, med, subMedicine)
			if err != nil {
				_ = tx.RollbackTx()
				return dto.Fail[any](response.CODE_FAIL_ADDITEM_PRESCRIPTION, err)
			}

			// 미리 복용내역 생성
			for i := 0; i <= req.TakeDays; i++ {
				futureTakeDate := p.mono.Date.TruncateToDateAddDay(startedAt, i)
				_, err = p.repo.AddTakeHistoryItemFromPrescriptionItemWithDate(pscItem, futureTakeDate)
				if err != nil {
					_ = tx.RollbackTx()
					return dto.Fail[any](response.CODE_FAIL_ADD_TAKE_HISTORY_ITEM, err)
				}
			}
		}
	}

	_ = tx.CommitTx()
	return dto.Ok[any](response.CODE_SUCCESS, nil)
}

// 복용계획 삭제
func (p *planService) Delete(req *DeletePlanRequest) dto.BaseResponse[any] {
	tx := p.repo.GetTxManager()
	err := tx.BeginTx()
	if err != nil {
		return dto.Fail[any](response.CODE_ERROR_TRANSACTION, err)
	}

	// 처방전 조회
	ps, err := p.repo.GetPrescription(req.ID)
	if err != nil {
		_ = tx.RollbackTx()
		return dto.Fail[any](response.CODE_NOT_FOUND_PLAN, err)
	}

	// 처방전 삭제
	err = p.repo.DeletePrescription(ps.ID)
	if err != nil {
		_ = tx.RollbackTx()
		return dto.Fail[any](response.CODE_FAIL_DELETE_PRESCRIPTION, err)
	}

	// 처방전 아이템 삭제
	err = p.repo.DeletePrescriptionItemByPrescriptionId(ps.ID)
	if err != nil {
		_ = tx.RollbackTx()
		return dto.Fail[any](response.CODE_FAIL_DELETEITEM_PRESCRIPTION, err)
	}

	return dto.Ok[any](response.CODE_SUCCESS, nil)
}

// 시간대 전체 복용처리
func (p *planService) Take(req *TakeToggleRequest) dto.BaseResponse[bool] {
	dateTime, err := p.mono.Date.ParseWithTime("Y-m-d H:i:s", req.TargetDate)
	if err != nil {
		return dto.Fail[bool](response.CODE_NOT_AVAILABLE_DATE, err)
	}

	// 트랜잭션 시작
	tx := p.repo.GetTxManager()
	err = tx.BeginTx()
	if err != nil {
		return dto.Fail[bool](response.CODE_ERROR_TRANSACTION, err)
	}

	// 날짜기준 유효한 처방전 조회
	psList, err := p.repo.GetPrescriptionInDateByTimeZoneId(req.UserId, dateTime, req.TimezoneId)
	if err != nil || psList == nil || len(psList) == 0 {
		_ = tx.RollbackTx()
		return dto.Fail[bool](response.CODE_NO_RESULT_PLAN_BY_DATE, err)
	}

	// 전체 의약품 수
	totalPills := 0
	// 복용된 의약품 수
	takenPills := 0

	// 처방전별 복용여부 확인
	for _, psc := range psList {
		items, err := psc.Edges.PrescriptionItemsOrErr()
		if err != nil {
			continue
		}

		// 전체 의약품 수 업데이트
		totalPills += len(items)

		// 의약품 복용여부 확인
		for _, item := range items {
			// 복용량
			takeAmount := item.TakeAmount
			// 남은량
			remainAmount := item.RemainAmount

			// 복용 내역이 존재하는지 확인
			tkhItem, err := p.repo.GetTakeHistoryItemFromPrescriptionItemByDate(req.UserId, item.ID, dateTime)
			if err != nil || tkhItem == nil {
				tkhItem, err = p.repo.AddTakeHistoryItemFromPrescriptionItemWithDate(item, dateTime)
				if err != nil {
					_ = tx.RollbackTx()
					return dto.Fail[bool](response.CODE_FAIL_ADD_TAKE_HISTORY_ITEM, err)
				}
			}

			// 복용량, 남은량을 복용내역 데이터로 치환
			takeAmount = tkhItem.TakeAmount
			remainAmount = tkhItem.RemainAmount

			// 복용량이 남은량보다 많은 경우
			if remainAmount < takeAmount {
				_ = tx.RollbackTx()
				return dto.Fail[bool](response.CODE_TOO_MUCH_TAKE_AMOUNT, err)
			}

			// 복용된 약인 경우 복용된 의약품 수 증가
			if tkhItem.TakeStatus {
				takenPills++
			}
		}
	}

	// 복용된 의약품 수가 전체 의약품 수와 같은 경우 복용처리 취소 플래그 설정
	takeFlag := takenPills < totalPills

	// 타임존에 대한 복용내역 조회
	hstItems, err := p.repo.GetTakeHistoryItemsByTimeZoneId(req.UserId, dateTime, req.TimezoneId)
	if err != nil {
		_ = tx.RollbackTx()
		return dto.Fail[bool](response.CODE_FAIL_GET_LIST_TAKE_HISTORY, err)
	}

	// 복용처리 - 복용내역 업데이트
	for _, item := range hstItems {
		// 기존 상태와 변경될 상태가 같은 경우엔 스킵
		if takeFlag == item.TakeStatus {
			continue
		}

		// 복용 상태 업데이트
		item.TakeStatus = takeFlag

		// 복용량 업데이트 (복용처리 취소인 경우 남은량 복구)
		if takeFlag {
			// 복용량이 남은량보다 많은 경우
			if item.RemainAmount < item.TakeAmount {
				_ = tx.RollbackTx()
				return dto.Fail[bool](response.CODE_TOO_MUCH_TAKE_AMOUNT, err)
			}
			item.RemainAmount = item.RemainAmount - item.TakeAmount
			item.TakeTime = p.mono.Date.GetTime(time.Now())
		} else {
			item.RemainAmount = item.RemainAmount + item.TakeAmount
			item.TakeTime = ""
		}

		// 복용내역 업데이트
		err = p.repo.UpdateTakeHistoryItem(item)
		if err != nil {
			_ = tx.RollbackTx()
			return dto.Fail[bool](response.CODE_FAIL_UPDATE_TAKEHISTORY, err)
		}

		// 처방전 아이템 업데이트
		err := p.repo.UpdatePrescriptionItemFromTakeHistoryItem(item)
		if err != nil {
			_ = tx.RollbackTx()
			return dto.Fail[bool](response.CODE_FAIL_UPDATE_PRESCRIPTION, err)
		}
	}

	// 트랜잭션 커밋
	err = tx.CommitTx()
	if err != nil {
		return dto.Fail[bool](response.CODE_FAIL_TAKE_PLAN, err)
	}
	return dto.Ok[bool](response.CODE_SUCCESS, &takeFlag)
}

// 개별 의약품 복용처리
func (p *planService) PillToggle(req *PillToggleRequest) dto.BaseResponse[bool] {
	// 트랜잭션 시작
	tx := p.repo.GetTxManager()
	err := tx.BeginTx()
	if err != nil {
		return dto.Fail[bool](response.CODE_ERROR_TRANSACTION, err)
	}

	// 처방전 아이템 조회
	psItem, err := p.repo.GetPrescriptionItem(req.PillId)
	if err != nil {
		_ = tx.RollbackTx()
		return dto.Fail[bool](response.CODE_NOT_FOUND_PLAN_ITEM, err)
	}

	// 복용내역이 없는 경우 복용내역 생성
	tkhItem, err := p.repo.GetTodayTakeHistoryItemFromPrescriptionItem(psItem.ID)
	if err != nil || tkhItem == nil {
		tkhItem, err = p.repo.AddTakeHistoryItemFromPrescriptionItem(psItem, time.Now())
		if err != nil {
			_ = tx.RollbackTx()
			return dto.Fail[bool](response.CODE_FAIL_ADD_TAKE_HISTORY_ITEM, err)
		}
	}

	tkhItem.TakeStatus = !tkhItem.TakeStatus
	if tkhItem.TakeStatus {
		// 복용량이 남은량보다 많은 경우
		if tkhItem.RemainAmount < tkhItem.TakeAmount {
			_ = tx.RollbackTx()
			return dto.Fail[bool](response.CODE_TOO_MUCH_TAKE_AMOUNT, err)
		}
		tkhItem.RemainAmount = tkhItem.RemainAmount - tkhItem.TakeAmount
		tkhItem.TakeDate = p.mono.Date.GetDate(time.Now())
		tkhItem.TakeTime = p.mono.Date.GetTime(time.Now())
	} else {
		tkhItem.RemainAmount = tkhItem.RemainAmount + tkhItem.TakeAmount
		tkhItem.TakeTime = ""
	}

	// 복용내역 업데이트
	err = p.repo.UpdateTakeHistoryItem(tkhItem)
	if err != nil {
		_ = tx.RollbackTx()
		return dto.Fail[bool](response.CODE_FAIL_UPDATE_TAKEHISTORY, err)
	}

	// 처방전 아이템 업데이트
	err = p.repo.UpdatePrescriptionItemFromTakeHistoryItem(tkhItem)
	if err != nil {
		_ = tx.RollbackTx()
		return dto.Fail[bool](response.CODE_FAIL_UPDATE_PRESCRIPTION, err)
	}

	_ = tx.CommitTx()
	return dto.Ok[bool](response.CODE_SUCCESS, &tkhItem.TakeStatus)
}

// 메모 업데이트
func (p *planService) UpdateMemo(req *UpdateMemoRequest) dto.BaseResponse[any] {
	dateTime, err := p.mono.Date.Parse("Y-m-d", req.Date)
	if err != nil {
		return dto.Fail[any](response.CODE_NOT_AVAILABLE_DATE, err)
	}

	// 메모 업데이트
	_ = p.repo.UpdateTakeHistoryMemo(req.UserId, dateTime, req.TimezoneId, req.Memo)

	return dto.Ok[any](response.CODE_SUCCESS, nil)
}

// 복용량 업데이트
func (p *planService) PillTakeAmountUpdate(req *PillTakeAmountUpdateRequest) dto.BaseResponse[bool] {
	// 복용량이 0 인경우 에러
	if req.TakeAmount <= 0 {
		return dto.Fail[bool](response.CODE_TOO_MUCH_TAKE_AMOUNT, nil)
	}

	// 트랜잭션 시작
	tx := p.repo.GetTxManager()
	err := tx.BeginTx()
	if err != nil {
		return dto.Fail[bool](response.CODE_ERROR_TRANSACTION, err)
	}

	// 처방전 아이템 조회
	psItem, err := p.repo.GetPrescriptionItem(req.PrescriptionItemId)
	if err != nil {
		_ = tx.RollbackTx()
		return dto.Fail[bool](response.CODE_NOT_FOUND_PLAN_ITEM, err)
	}

	// 복용내역이 없는 경우 복용내역 생성
	tkhItem, err := p.repo.GetTodayTakeHistoryItemFromPrescriptionItem(psItem.ID)
	if err != nil || tkhItem == nil {
		tkhItem, err = p.repo.AddTakeHistoryItemFromPrescriptionItem(psItem, time.Now())
		if err != nil {
			_ = tx.RollbackTx()
			return dto.Fail[bool](response.CODE_FAIL_ADD_TAKE_HISTORY_ITEM, err)
		}
	}

	// 원본 복용량
	originTakeAmount := tkhItem.TakeAmount
	// 변경될 복용량
	tkhItem.TakeAmount = req.TakeAmount

	// 현재 복용상태인 경우
	if tkhItem.TakeStatus {
		// 이미 복용된 상태라면 남은량 싱크를 맞춰준다.
		tkhItem.RemainAmount = (tkhItem.RemainAmount + originTakeAmount) - tkhItem.TakeAmount
		// 합산 수치가 정합성이 맞지않는 경우
		if tkhItem.RemainAmount < 0 || tkhItem.RemainAmount > tkhItem.TotalAmount {
			_ = tx.RollbackTx()
			return dto.Fail[bool](response.CODE_NOT_VALID_AMOUNT, err)
		}
		// 남은량이 음수인 경우
		if tkhItem.RemainAmount < 0 {
			_ = tx.RollbackTx()
			return dto.Fail[bool](response.CODE_REMAIN_AMOUNT_CANNOT_BE_NEGATIVE, err)
		}
	} else {
		// 복용량이 남은량보다 많은 경우
		if tkhItem.RemainAmount < tkhItem.TakeAmount {
			_ = tx.RollbackTx()
			return dto.Fail[bool](response.CODE_TOO_MUCH_TAKE_AMOUNT, err)
		}
	}

	// 복용량 업데이트
	err = p.repo.UpdateTakeHistoryItem(tkhItem)
	if err != nil {
		_ = tx.RollbackTx()
		return dto.Fail[bool](response.CODE_FAIL_UPDATE_TAKEHISTORY, err)
	}

	// 처방전 아이템 업데이트
	err = p.repo.UpdatePrescriptionItemFromTakeHistoryItem(tkhItem)
	if err != nil {
		_ = tx.RollbackTx()
		return dto.Fail[bool](response.CODE_FAIL_UPDATE_PRESCRIPTION, err)
	}

	// 커밋 처리
	_ = tx.CommitTx()
	return dto.Ok[bool](response.CODE_SUCCESS, nil)
}
