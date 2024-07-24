package plan

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"nursing_api/pkg/database"
	"nursing_api/pkg/ent"
	mdSchema "nursing_api/pkg/ent/medicine"
	pscSchema "nursing_api/pkg/ent/prescription"
	pscItmSchema "nursing_api/pkg/ent/prescriptionitem"
	tkhItmSchema "nursing_api/pkg/ent/takehistoryitem"
	tkhMemoSchema "nursing_api/pkg/ent/takehistorymemo"
	tzSchema "nursing_api/pkg/ent/timezone"
	"nursing_api/pkg/mono"
	"strconv"
	"time"
)

type Repository interface {
	GetTxManager() database.TransactionManager

	GetPrescription(id uuid.UUID) (*ent.Prescription, error)
	GetPrescriptionInDate(userId uuid.UUID, date time.Time) ([]*ent.Prescription, error)
	GetPrescriptionInDateByTimeZoneId(userId uuid.UUID, date time.Time, timezoneId uuid.UUID) ([]*ent.Prescription, error)
	GetPrescriptionInMonth(userId uuid.UUID, year int, month int) ([]*ent.Prescription, error)
	GetPrescriptionItem(id uuid.UUID) (*ent.PrescriptionItem, error)
	GetTimeZone(id uuid.UUID) (*ent.TimeZone, error)
	GetTakeHistoryItem(takeHistoryItemId uuid.UUID) (*ent.TakeHistoryItem, error)
	GetTodayTakeHistoryItemFromPrescriptionItem(prescriptionItemId uuid.UUID) (*ent.TakeHistoryItem, error)
	GetTakeHistoryItemFromPrescriptionItemByDate(userId uuid.UUID, prescriptionItemId uuid.UUID, date time.Time) (*ent.TakeHistoryItem, error)
	GetMedicine(id uuid.UUID) (*ent.Medicine, error)

	AddPrescription(plan AddPlanRequest) (*ent.Prescription, error)
	AddPrescriptionItems(prescription *ent.Prescription, tz *ent.TimeZone, medicine *ent.Medicine, subMedicine AddMedicineRequest) (*ent.PrescriptionItem, error)
	AddTakeHistoryItemFromPrescriptionItem(prescriptionItem *ent.PrescriptionItem, date time.Time) (*ent.TakeHistoryItem, error)
	AddTakeHistoryItemFromPrescriptionItemWithDate(prescriptionItem *ent.PrescriptionItem, date time.Time) (*ent.TakeHistoryItem, error)
	UpdateTakeHistoryItem(takeHistoryItem *ent.TakeHistoryItem) error
	UpdatePrescriptionItemFromTakeHistoryItem(takeHistoryItem *ent.TakeHistoryItem) error
	UpdateTakeHistoryMemo(userId uuid.UUID, date time.Time, timezoneId uuid.UUID, memo string) error
	DeletePrescription(id uuid.UUID) error
	DeletePrescriptionItemByPrescriptionId(prescriptionId uuid.UUID) error

	GetTakeHistoryItemsByTimeZoneId(userId uuid.UUID, date time.Time, timezoneId uuid.UUID) ([]*ent.TakeHistoryItem, error)
	GetTakeHistoryMemoByTimeZoneId(userId uuid.UUID, date time.Time, timezoneId uuid.UUID) (*ent.TakeHistoryMemo, error)

	GetPlanMap(userId uuid.UUID, date time.Time) (map[uuid.UUID]*Plan, error)
}

type repository struct {
	db   database.DatabaseClient
	mono *mono.Client
}

func NewRepository(
	db database.DatabaseClient,
	mono *mono.Client,
) Repository {
	return &repository{
		db:   db,
		mono: mono,
	}
}

// 처방전 조회
func (r repository) GetPrescription(id uuid.UUID) (*ent.Prescription, error) {
	ps, err := r.prescriptionClient().
		Query().
		Where(pscSchema.IDEQ(id)).
		Only(r.GetCtx())
	if err != nil {
		return nil, err
	}

	return ps, nil
}

// 해당일자 유효한 처방전 조회
func (r repository) GetPrescriptionInDate(userId uuid.UUID, date time.Time) ([]*ent.Prescription, error) {
	ps, err := r.prescriptionClient().
		Query().
		WithPrescriptionItems(func(query *ent.PrescriptionItemQuery) {
			query.Order(pscItmSchema.ByCreatedAt())
		}).
		Where(
			pscSchema.And(
				pscSchema.UserID(userId),
				pscSchema.StartedAtLTE(r.mono.Date.TruncateToDate(date)),
				pscSchema.FinishedAtGTE(r.mono.Date.TruncateToDate(date)),
			),
		).
		All(r.GetCtx())
	if err != nil {
		return nil, err
	}

	return ps, nil
}

// 해당년도 월별 유효한 처방전 조회
func (r repository) GetPrescriptionInMonth(userId uuid.UUID, year int, month int) ([]*ent.Prescription, error) {
	found, err := r.prescriptionClient().
		Query().
		Where(
			pscSchema.UserIDEQ(userId),
			pscSchema.StartedAtLT(r.mono.Date.ToFirstDayOfMonth(year, month+1)),
			pscSchema.FinishedAtGTE(r.mono.Date.ToFirstDayOfMonth(year, month)),
		).
		All(r.GetCtx())
	if err != nil {
		return nil, err
	}

	return found, nil
}

// 해당일자 타임존별 유효한 처방전 조회
func (r repository) GetPrescriptionInDateByTimeZoneId(
	userId uuid.UUID,
	date time.Time,
	timezoneId uuid.UUID,
) ([]*ent.Prescription, error) {
	ps, err := r.prescriptionClient().
		Query().
		WithPrescriptionItems(func(query *ent.PrescriptionItemQuery) {
			query.WithTakeHistoryItem()
			query.Where(pscItmSchema.TimezoneIDEQ(timezoneId))
		}).
		Where(
			pscSchema.And(
				pscSchema.UserID(userId),
				pscSchema.StartedAtLTE(r.mono.Date.TruncateToDate(date)),
				pscSchema.FinishedAtGT(r.mono.Date.TruncateToDateAddDay(date, 1)),
			),
		).
		All(r.GetCtx())
	if err != nil {
		return nil, err
	}

	return ps, nil
}

// 처방전 아이템 조회
func (r repository) GetPrescriptionItem(id uuid.UUID) (*ent.PrescriptionItem, error) {
	pscItm, err := r.prescriptionItemClient().
		Query().
		Where(pscItmSchema.IDEQ(id)).
		Only(r.GetCtx())
	if err != nil {
		return nil, err
	}

	return pscItm, nil
}

// 타임존 조회
func (r repository) GetTimeZone(id uuid.UUID) (*ent.TimeZone, error) {
	tz, err := r.timeZoneClient().
		Query().
		Where(tzSchema.IDEQ(id)).
		Only(r.GetCtx())
	if err != nil {
		return nil, err
	}

	return tz, nil
}

// 복용내역 조회
func (r repository) GetTakeHistoryItem(takeHistoryItemId uuid.UUID) (*ent.TakeHistoryItem, error) {
	tkhItm, err := r.takeHistoryItemClient().
		Query().
		Where(tkhItmSchema.IDEQ(takeHistoryItemId)).
		Only(r.GetCtx())
	if err != nil {
		return nil, err
	}

	return tkhItm, nil
}

// 오늘 개별 의약품 복용내역 조회
func (r repository) GetTodayTakeHistoryItemFromPrescriptionItem(
	prescriptionItemId uuid.UUID,
) (*ent.TakeHistoryItem, error) {
	tkhItm, err := r.takeHistoryItemClient().
		Query().
		Where(
			tkhItmSchema.And(
				tkhItmSchema.TakeDateEQ(r.mono.Date.GetDate(time.Now())),
				tkhItmSchema.PrescriptionItemIDEQ(prescriptionItemId),
			),
		).
		Only(r.GetCtx())
	if err != nil {
		return nil, err
	}

	return tkhItm, nil
}

// 날짜별 개별 의약품 복용내역 조회
func (r repository) GetTakeHistoryItemFromPrescriptionItemByDate(
	userId uuid.UUID,
	prescriptionItemId uuid.UUID,
	date time.Time,
) (*ent.TakeHistoryItem, error) {
	tkhItm, err := r.takeHistoryItemClient().
		Query().
		Where(
			tkhItmSchema.And(
				tkhItmSchema.UserIDEQ(userId),
				tkhItmSchema.TakeDateEQ(r.mono.Date.GetDate(date)),
				tkhItmSchema.PrescriptionItemIDEQ(prescriptionItemId),
			),
		).
		Limit(1).
		Only(r.GetCtx())
	if err != nil {
		return nil, err
	}

	return tkhItm, nil
}

// 의약품 조회
func (r repository) GetMedicine(id uuid.UUID) (*ent.Medicine, error) {
	med, err := r.medicineClient().
		Query().
		Where(mdSchema.IDEQ(id)).
		Only(r.GetCtx())
	if err != nil {
		return nil, err
	}

	return med, nil
}

// 처방전 생성
func (r repository) AddPrescription(plan AddPlanRequest) (*ent.Prescription, error) {
	startedAt := r.mono.Date.ParseForce("Y-m-d", plan.StartedAt)
	finishedAt := startedAt.AddDate(0, 0, plan.TakeDays)
	ps, err := r.prescriptionClient().
		Create().
		SetUserID(plan.UserId).
		SetPrescriptionName(plan.Name).
		SetHospitalName(plan.Hospital).
		SetTakeDays(plan.TakeDays).
		SetStartedAt(startedAt).
		SetFinishedAt(finishedAt).
		Save(r.GetCtx())
	if err != nil {
		return nil, err
	}

	return ps, nil
}

// 처방전 아이템 생성
func (r repository) AddPrescriptionItems(
	prescription *ent.Prescription,
	tz *ent.TimeZone,
	medicine *ent.Medicine,
	subMedicine AddMedicineRequest,
) (*ent.PrescriptionItem, error) {
	pscItm, err := r.prescriptionItemClient().
		Create().
		SetUserID(prescription.UserID).
		SetPrescriptionID(prescription.ID).
		SetTimezoneID(tz.ID).
		SetTimezoneName(tz.TimezoneName).
		SetMidday(tz.Midday).
		SetHour(tz.Hour).
		SetMinute(tz.Minute).
		SetMedicineID(medicine.ID).
		SetMedicineName(medicine.MedicineName).
		SetTakeAmount(subMedicine.TakeAmount).
		SetRemainAmount(subMedicine.RemainAmount).
		SetTotalAmount(subMedicine.TotalAmount).
		SetMedicineUnit(subMedicine.TakeUnit).
		SetMemo(subMedicine.Memo).
		Save(r.GetCtx())
	if err != nil {
		return nil, err
	}

	return pscItm, nil
}

// 처방전 아이템 복용내역 생성
func (r repository) AddTakeHistoryItemFromPrescriptionItem(
	prescriptionItem *ent.PrescriptionItem,
	date time.Time,
) (*ent.TakeHistoryItem, error) {
	tkhItm, err := r.takeHistoryItemClient().
		Create().
		SetUserID(prescriptionItem.UserID).
		SetPrescriptionItemID(prescriptionItem.ID).
		SetPrescriptionID(prescriptionItem.PrescriptionID).
		SetTimezoneID(prescriptionItem.TimezoneID).
		SetTimezoneName(prescriptionItem.TimezoneName).
		SetHour(prescriptionItem.Hour).
		SetMinute(prescriptionItem.Minute).
		SetMidday(prescriptionItem.Midday).
		SetMedicineID(prescriptionItem.MedicineID).
		SetMedicineName(prescriptionItem.MedicineName).
		SetTakeAmount(prescriptionItem.TakeAmount).
		SetRemainAmount(prescriptionItem.RemainAmount).
		SetTotalAmount(prescriptionItem.TotalAmount).
		SetTakeUnit(prescriptionItem.MedicineUnit).
		SetTakeDate(r.mono.Date.GetDate(date)).
		SetTakeTime("").
		SetTakeStatus(false).
		Save(r.GetCtx())
	if err != nil {
		return nil, err
	}

	return tkhItm, nil
}

// 처방전 아이템 복용내역 생성
func (r repository) AddTakeHistoryItemFromPrescriptionItemWithDate(
	prescriptionItem *ent.PrescriptionItem,
	date time.Time,
) (*ent.TakeHistoryItem, error) {
	tkhItm, err := r.takeHistoryItemClient().
		Create().
		SetUserID(prescriptionItem.UserID).
		SetPrescriptionItemID(prescriptionItem.ID).
		SetPrescriptionID(prescriptionItem.PrescriptionID).
		SetTimezoneID(prescriptionItem.TimezoneID).
		SetTimezoneName(prescriptionItem.TimezoneName).
		SetHour(prescriptionItem.Hour).
		SetMinute(prescriptionItem.Minute).
		SetMidday(prescriptionItem.Midday).
		SetMedicineID(prescriptionItem.MedicineID).
		SetMedicineName(prescriptionItem.MedicineName).
		SetTakeAmount(prescriptionItem.TakeAmount).
		SetRemainAmount(prescriptionItem.RemainAmount).
		SetTotalAmount(prescriptionItem.TotalAmount).
		SetTakeUnit(prescriptionItem.MedicineUnit).
		SetTakeDate(r.mono.Date.GetDate(date)).
		SetTakeTime("").
		SetTakeStatus(false).
		Save(r.GetCtx())
	if err != nil {
		return nil, err
	}

	return tkhItm, nil
}

// 복용내역 업데이트
func (r repository) UpdateTakeHistoryItem(takeHistoryItem *ent.TakeHistoryItem) error {
	_, err := r.takeHistoryItemClient().
		UpdateOne(takeHistoryItem).
		SetTakeDate(takeHistoryItem.TakeDate).
		SetTakeTime(takeHistoryItem.TakeTime).
		SetTakeStatus(takeHistoryItem.TakeStatus).
		SetTakeAmount(takeHistoryItem.TakeAmount).
		SetRemainAmount(takeHistoryItem.RemainAmount).
		SetUpdatedAt(time.Now()).
		Save(r.GetCtx())
	if err != nil {
		return err
	}

	return nil
}

// 처방전 아이템 업데이트 (복용내역에서 업데이트)
func (r repository) UpdatePrescriptionItemFromTakeHistoryItem(takeHistoryItem *ent.TakeHistoryItem) error {
	_, err := r.prescriptionItemClient().
		Update().
		SetRemainAmount(takeHistoryItem.RemainAmount).
		SetUpdatedAt(time.Now()).
		Where(
			pscItmSchema.IDEQ(takeHistoryItem.PrescriptionItemID),
		).
		Save(r.GetCtx())
	if err != nil {
		return err
	}

	return nil
}

// 복용메모 업데이트
func (r repository) UpdateTakeHistoryMemo(
	userId uuid.UUID,
	date time.Time,
	timezoneId uuid.UUID,
	memo string,
) error {
	tkhMemo, err := r.takeHistoryMemoClient().
		Query().
		Where(
			tkhMemoSchema.UserIDEQ(userId),
			tkhMemoSchema.TakeDateEQ(r.mono.Date.GetDate(date)),
			tkhMemoSchema.TimezoneIDEQ(timezoneId),
		).
		Only(r.GetCtx())
	if err != nil {
		_, err = r.takeHistoryMemoClient().
			Create().
			SetUserID(userId).
			SetTimezoneID(timezoneId).
			SetTakeDate(r.mono.Date.GetDate(date)).
			SetMemo(memo).
			Save(r.GetCtx())
		if err != nil {
			return err
		}
	} else {
		_, err = r.takeHistoryMemoClient().
			UpdateOne(tkhMemo).
			SetMemo(memo).
			Save(r.GetCtx())
		if err != nil {
			return err
		}
	}

	return nil
}

// 처방전 삭제
func (r repository) DeletePrescription(id uuid.UUID) error {
	_, err := r.prescriptionClient().Delete().Where(pscSchema.IDEQ(id)).Exec(r.GetCtx())
	if err != nil {
		return err
	}

	// 에러가 없는 경우 삭제 카운트가 없어도 성공으로 처리
	return nil
}

// 처방전 아이템 삭제
func (r repository) DeletePrescriptionItemByPrescriptionId(prescriptionId uuid.UUID) error {
	_, err := r.prescriptionItemClient().Delete().Where(pscItmSchema.PrescriptionIDEQ(prescriptionId)).Exec(r.GetCtx())
	if err != nil {
		return err
	}

	// 에러가 없는 경우 삭제 카운트가 없어도 성공으로 처리
	return nil
}

// 복용내역 조회
func (r repository) GetTakeHistoryItemsByTimeZoneId(userId uuid.UUID, date time.Time, timezoneId uuid.UUID) ([]*ent.TakeHistoryItem, error) {
	ths, err := r.takeHistoryItemClient().
		Query().
		Where(
			tkhItmSchema.UserIDEQ(userId),
			tkhItmSchema.TakeDateEQ(r.mono.Date.GetDate(date)),
			tkhItmSchema.TimezoneIDEQ(timezoneId),
		).
		All(r.GetCtx())
	if err != nil {
		return nil, err
	}

	return ths, nil
}

// 복용메모 조회
func (r repository) GetTakeHistoryMemoByTimeZoneId(userId uuid.UUID, date time.Time, timezoneId uuid.UUID) (*ent.TakeHistoryMemo, error) {
	thm, err := r.takeHistoryMemoClient().
		Query().
		Where(
			tkhMemoSchema.UserIDEQ(userId),
			tkhMemoSchema.TimezoneIDEQ(timezoneId),
			tkhMemoSchema.TakeDateEQ(r.mono.Date.GetDate(date)),
		).
		Only(r.GetCtx())
	if err != nil {
		return nil, err
	}

	return thm, nil
}

// 복용계획 조회
func (r repository) GetPlanMap(userId uuid.UUID, date time.Time) (map[uuid.UUID]*Plan, error) {
	// 현재날짜에 유효한 처방전 조회
	pscList, err := r.GetPrescriptionInDate(userId, date)
	if err != nil || pscList == nil || len(pscList) == 0 {
		return nil, err
	}

	// 처방전 아이템을 시간대별 복용계획 형태로 가공
	tzMap := map[uuid.UUID]*Plan{}
	for _, psc := range pscList {
		pscItems, err := psc.Edges.PrescriptionItemsOrErr()
		if err != nil {
			continue
		}

		// 처방전 아이템을 시간대별로 그룹핑
		for _, pscItm := range pscItems {
			tzId := pscItm.TimezoneID

			// 시간 파싱
			h, err := strconv.Atoi(pscItm.Hour)
			if err != nil {
				return nil, err
			}

			// 오후 시간대면서 12 이하인 경우 시간을 12시간 추가
			if h < 12 && pscItm.Midday == "PM" {
				pscItm.Hour = strconv.Itoa(h + 12)
			}

			// 타임존이 없는 경우 생성
			if tzMap[tzId] == nil {
				tzText := fmt.Sprintf("%s %s:%s", pscItm.Midday, pscItm.Hour, pscItm.Minute)
				hour, _ := strconv.Atoi(pscItm.Hour)
				minute, _ := strconv.Atoi(pscItm.Minute)
				tzMap[tzId] = &Plan{
					TimezoneId: tzId,
					PlanName:   pscItm.TimezoneName,
					Hour:       hour,
					Minute:     minute,
					Timezone:   tzText,
					Memo:       "",
					TakeStatus: false,
				}
			}

			// 복용내역이 있는 경우 업데이트
			tkhItem, err := r.GetTakeHistoryItemFromPrescriptionItemByDate(userId, pscItm.ID, date)
			takeHistoryItemId := uuid.Nil
			takeDate := ""
			takeTime := ""
			takeStatus := false
			takeAmount := pscItm.TakeAmount
			if err == nil {
				takeHistoryItemId = tkhItem.ID
				takeStatus = tkhItem.TakeStatus
				takeDate = tkhItem.TakeDate
				takeTime = tkhItem.TakeTime
				takeAmount = tkhItem.TakeAmount
			}

			// 처방전 아이템을 복용계획으로 추가
			newPill := Pill{
				PrescriptionId:     pscItm.PrescriptionID,
				PrescriptionItemId: pscItm.ID,
				MedicineId:         pscItm.MedicineID,
				PillName:           pscItm.MedicineName,
				TakeUnit:           pscItm.MedicineUnit,
				TakeAmount:         takeAmount,
				RemainAmount:       pscItm.RemainAmount,
				TotalAmount:        pscItm.TotalAmount,
				TakeHistoryItemId:  takeHistoryItemId,
				TakeDate:           takeDate,
				TakeTime:           takeTime,
				TakeStatus:         takeStatus,
			}
			tzMap[tzId].Pills = append(tzMap[tzId].Pills, newPill)
		}
	}

	// 시간대별 메모 조회 후 복용계획에 추가
	for tzId, plan := range tzMap {
		// 복용내역이 모두 완료된 경우 복용완료로 변경
		pillTotal := len(plan.Pills)
		taken := 0

		// 복용완료 시간 중 가장 늦은 시간을 기준으로 복용완료 처리
		var latestTime time.Time
		for i := 0; i < pillTotal; i++ {
			if plan.Pills[i].TakeStatus {
				taken++
			}

			// 복용완료 시간 중 가장 늦은 시간을 기준으로 복용완료 처리
			if plan.Pills[i].TakeTime != "" {
				takeTime, err := time.Parse("15:04:05", plan.Pills[i].TakeTime)
				if err == nil && latestTime.After(takeTime) {
					latestTime = takeTime
				}
			}
		}

		if pillTotal == taken {
			plan.TakeStatus = true
			plan.TakeTime = latestTime.Format("15:04:05")
		}

		tkhMemo, err := r.GetTakeHistoryMemoByTimeZoneId(userId, date, tzId)
		if err != nil {
			continue
		}
		plan.Memo = tkhMemo.Memo
		tzMap[tzId] = plan
	}

	return tzMap, nil
}

func (r repository) timeZoneClient() *ent.TimeZoneClient {
	return r.db.GetClient().TimeZone
}

func (r repository) prescriptionClient() *ent.PrescriptionClient {
	return r.db.GetClient().Debug().Prescription
}

func (r repository) prescriptionItemClient() *ent.PrescriptionItemClient {
	return r.db.GetClient().Debug().PrescriptionItem
}

func (r repository) takeHistoryItemClient() *ent.TakeHistoryItemClient {
	return r.db.GetClient().Debug().TakeHistoryItem
}

func (r repository) takeHistoryMemoClient() *ent.TakeHistoryMemoClient {
	return r.db.GetClient().Debug().TakeHistoryMemo
}

func (r repository) medicineClient() *ent.MedicineClient {
	return r.db.GetClient().Medicine
}

func (r repository) GetTxManager() database.TransactionManager {
	return r.db.GetTxManager()
}

func (r repository) GetCtx() context.Context {
	return r.db.GetCtx()
}
