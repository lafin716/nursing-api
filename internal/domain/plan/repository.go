package plan

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"nursing_api/pkg/database"
	"nursing_api/pkg/ent"
	pscSchema "nursing_api/pkg/ent/prescription"
	pscItmSchema "nursing_api/pkg/ent/prescriptionitem"
	tkhItmSchema "nursing_api/pkg/ent/takehistoryitem"
	tkhMemoSchema "nursing_api/pkg/ent/takehistorymemo"
	tzSchema "nursing_api/pkg/ent/timezone"
	"nursing_api/pkg/mono"
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

	AddPrescription(plan AddPlanRequest) (*ent.Prescription, error)
	AddPrescriptionItems(prescriptionId uuid.UUID, tz *ent.TimeZone, medicine AddMedicineRequest) (*ent.PrescriptionItem, error)
	AddTakeHistoryItemFromPrescriptionItem(prescriptionItem *ent.PrescriptionItem) (*ent.TakeHistoryItem, error)
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
		WithPrescriptionItems().
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

// 해당년도 월별 유효한 처방전 조회
func (r repository) GetPrescriptionInMonth(userId uuid.UUID, year int, month int) ([]*ent.Prescription, error) {
	found, err := r.prescriptionClient().
		Query().
		Where(
			pscSchema.UserIDEQ(userId),
			pscSchema.StartedAtGTE(time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Now().Location())),
			pscSchema.StartedAtLT(time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.Now().Location())),
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
		Where(pscItmSchema.PrescriptionIDEQ(id)).
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
	prescriptionId uuid.UUID,
	tz *ent.TimeZone,
	medicine AddMedicineRequest,
) (*ent.PrescriptionItem, error) {
	pscItm, err := r.prescriptionItemClient().
		Create().
		SetPrescriptionID(prescriptionId).
		SetTimezoneID(tz.ID).
		SetTimezoneName(tz.TimezoneName).
		SetMidday(tz.Midday).
		SetHour(tz.Hour).
		SetMinute(tz.Minute).
		SetMedicineID(medicine.MedicineId).
		SetMedicineName(medicine.Name).
		SetTakeAmount(medicine.TakeAmount).
		SetRemainAmount(medicine.RemainAmount).
		SetTotalAmount(medicine.TotalAmount).
		SetMedicineUnit(medicine.TakeUnit).
		SetMemo(medicine.Memo).
		Save(r.GetCtx())
	if err != nil {
		return nil, err
	}

	return pscItm, nil
}

// 처방전 아이템 복용내역 생성
func (r repository) AddTakeHistoryItemFromPrescriptionItem(
	prescriptionItem *ent.PrescriptionItem,
) (*ent.TakeHistoryItem, error) {
	tkhItm, err := r.takeHistoryItemClient().
		Create().
		SetUserID(prescriptionItem.Edges.Prescription.UserID).
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
			tkhMemoSchema.TakeDateEQ(date),
			tkhMemoSchema.TimezoneIDEQ(timezoneId),
		).
		Only(r.GetCtx())
	if err != nil {
		_, err = r.takeHistoryMemoClient().
			Create().
			SetUserID(userId).
			SetTimezoneID(timezoneId).
			SetTakeDate(date).
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
			tkhItmSchema.TakeDateEQ(date),
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
			tkhMemoSchema.TakeDateEQ(date),
			tkhMemoSchema.TimezoneIDEQ(timezoneId),
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
		pscItms, err := psc.Edges.PrescriptionItemsOrErr()
		if err != nil {
			continue
		}

		// 처방전 아이템을 시간대별로 그룹핑
		for _, pscItm := range pscItms {
			tzId := pscItm.TimezoneID

			// 타임존이 없는 경우 생성
			if tzMap[tzId] == nil {
				tzText := fmt.Sprintf("%s %s:%s", pscItm.Midday, pscItm.Hour, pscItm.Minute)
				tzMap[tzId] = &Plan{
					TimezoneId: tzId,
					PlanName:   pscItm.TimezoneName,
					Timezone:   tzText,
					Memo:       "",
				}
			}

			// 복용내역이 있는 경우 업데이트
			tkhItm, err := pscItm.Edges.TakeHistoryItemOrErr()
			takeHistoryItemId := uuid.Nil
			takeDate := ""
			takeStatus := false
			if err == nil {
				takeHistoryItemId = tkhItm.ID
				takeStatus = tkhItm.TakeStatus
				takeDate = tkhItm.TakeDate.Format("2006-01-02 15:04:05")
			}

			// 처방전 아이템을 복용계획으로 추가
			newPill := Pill{
				PrescriptionId:     pscItm.PrescriptionID,
				PrescriptionItemId: pscItm.ID,
				MedicineId:         pscItm.MedicineID,
				PillName:           pscItm.MedicineName,
				TakeUnit:           pscItm.MedicineUnit,
				TakeAmount:         pscItm.TakeAmount,
				RemainAmount:       pscItm.RemainAmount,
				TotalAmount:        pscItm.TotalAmount,
				TakeHistoryItemId:  takeHistoryItemId,
				TakeDate:           takeDate,
				TakeStatus:         takeStatus,
			}
			tzMap[tzId].Pills = append(tzMap[tzId].Pills, newPill)
		}
	}

	// 시간대별 메모 조회 후 복용계획에 추가
	for tzId, plan := range tzMap {
		tkhMemo, err := r.GetTakeHistoryMemoByTimeZoneId(userId, date, tzId)
		if err != nil {
			continue
		}
		plan.Memo = tkhMemo.Memo
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
	return r.db.GetClient().PrescriptionItem
}

func (r repository) takeHistoryItemClient() *ent.TakeHistoryItemClient {
	return r.db.GetClient().TakeHistoryItem
}

func (r repository) takeHistoryMemoClient() *ent.TakeHistoryMemoClient {
	return r.db.GetClient().TakeHistoryMemo
}

func (r repository) GetTxManager() database.TransactionManager {
	return r.db.GetTxManager()
}

func (r repository) GetCtx() context.Context {
	return r.db.GetCtx()
}

func Convert[T any, O any](data []O, action func(O) T) []T {
	var result []T
	for _, d := range data {
		result = append(result, action(d))
	}

	return result
}
