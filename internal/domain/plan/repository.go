package plan

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	takehistory "nursing_api/internal/domain/take_history"
	"nursing_api/pkg/database"
	"nursing_api/pkg/ent"
	scPs "nursing_api/pkg/ent/prescription"
	scPsItem "nursing_api/pkg/ent/prescriptionitem"
	scTh "nursing_api/pkg/ent/takehistory"
	scThItem "nursing_api/pkg/ent/takehistoryitem"
	scTzl "nursing_api/pkg/ent/timezonelink"
	"nursing_api/pkg/mono"
	"sort"
	"strconv"
	"time"
)

type Repository interface {
	GetPlans(userId uuid.UUID, date time.Time) ([]Plan, error)
	GetPlansByMonth(userId uuid.UUID, year int, month int) (*Summary, error)
}

type repository struct {
	root             *ent.Client
	mono             *mono.Client
	prescription     *ent.PrescriptionClient
	prescriptionItem *ent.PrescriptionItemClient
	timezoneLink     *ent.TimeZoneLinkClient
	takeHistory      *ent.TakeHistoryClient
	takeHistoryItem  *ent.TakeHistoryItemClient
	ctx              context.Context
}

func NewRepository(
	mono *mono.Client,
	db *database.DatabaseClient,
) Repository {
	return &repository{
		root:             db.Client,
		mono:             mono,
		prescription:     db.Client.Prescription,
		prescriptionItem: db.Client.PrescriptionItem,
		timezoneLink:     db.Client.TimeZoneLink,
		takeHistory:      db.Client.TakeHistory,
		takeHistoryItem:  db.Client.TakeHistoryItem,
		ctx:              db.Ctx,
	}
}

func (r repository) GetPlans(userId uuid.UUID, date time.Time) ([]Plan, error) {
	// 현재날짜에 유효한 처방전 조회
	pcs, err := r.prescription.
		Query().
		Where(
			scPs.And(
				scPs.UserID(userId),
				scPs.StartedAtLTE(date),
				scPs.FinishedAtGTE(date),
			),
		).All(r.ctx)
	if err != nil {
		return []Plan{}, err
	}

	if len(pcs) == 0 {
		return []Plan{}, nil
	}

	// 처방전 고유번호 슬라이스
	pcsIds := Convert[uuid.UUID, *ent.Prescription](pcs, func(data *ent.Prescription) uuid.UUID { return data.ID })

	// 처방전과 연결된 타임존 조회
	tzls, err := r.timezoneLink.
		Query().
		Where(
			scTzl.PrescriptionIDIn(pcsIds...),
		).
		Order(
			ent.Asc(scTzl.FieldHour),
			ent.Asc(scTzl.FieldMinute),
		).
		All(r.ctx)

	// 타임존 고유번호 슬라이스
	tzIds := Convert[uuid.UUID, *ent.TimeZoneLink](tzls, func(data *ent.TimeZoneLink) uuid.UUID { return data.TimezoneID })

	// 타임존 고유번호로 타임존링크 고유번호 그룹핑
	tzMap := map[uuid.UUID]*Plan{}
	for _, tzl := range tzls {
		pcsItems, err := r.prescriptionItem.
			Query().
			Where(scPsItem.TimezoneLinkID(tzl.ID)).
			All(r.ctx)
		if err != nil {
			return []Plan{}, err
		}

		var pills []Pill
		for _, pcsItem := range pcsItems {
			pills = append(pills, Pill{
				PrescriptionItemId: pcsItem.ID,
				PillName:           pcsItem.MedicineName,
				MedicineId:         pcsItem.MedicineID,
				TakeUnit:           pcsItem.MedicineUnit,
				TakeAmount:         pcsItem.TakeAmount,
				RemainAmount:       pcsItem.RemainAmount,
				TotalAmount:        pcsItem.TotalAmount,
			})
		}

		var plan *Plan
		if tzMap[tzl.TimezoneID] == nil {
			tzText := fmt.Sprintf("%s %s:%s", tzl.Midday, tzl.Hour, tzl.Minute)
			h, err := strconv.Atoi(tzl.Hour)
			if err != nil {
				return []Plan{}, err
			}
			m, err := strconv.Atoi(tzl.Minute)
			if err != nil {
				return []Plan{}, err
			}
			plan = &Plan{
				TimezoneId: tzl.TimezoneID,
				PlanName:   tzl.TimezoneName,
				Hour:       h,
				Minute:     m,
				Timezone:   tzText,
			}
		} else {
			plan = tzMap[tzl.TimezoneID]
		}

		plan.Pills = append(plan.Pills, pills...)
		tzMap[tzl.TimezoneID] = plan
	}

	// 복용내역 조회
	histories, err := r.takeHistory.
		Query().
		Where(
			scTh.And(
				scTh.UserID(userId),
				scTh.TimezoneIDIn(tzIds...),
				scTh.TakeDateGTE(date),
				scTh.TakeDateLT(date.AddDate(0, 0, 1)),
			),
		).All(r.ctx)
	if err != nil {
		return []Plan{}, err
	}

	// 복용여부 맵
	takeStatusMap := map[uuid.UUID]*Plan{}
	for _, history := range histories {
		takeStatusMap[history.TimezoneID] = &Plan{
			TakeStatus: takehistory.DONE == takehistory.TakeStatus(history.TakeStatus),
			TakeDate:   r.mono.Date.Format(history.TakeDate, "Y-m-d H:i:s"),
			Memo:       history.Memo,
		}
	}

	// 복용내역 고유번호 슬라이스
	historyIds := Convert[uuid.UUID, *ent.TakeHistory](histories, func(data *ent.TakeHistory) uuid.UUID { return data.ID })

	// 복용의약품내역 조회
	historyItems, err := r.takeHistoryItem.
		Query().
		Where(
			scThItem.TakeHistoryIDIn(historyIds...),
		).
		All(r.ctx)
	if err != nil {
		return []Plan{}, err
	}

	// 의약품 복용여부 맵
	takePillStatusMap := map[uuid.UUID]*Pill{}
	for _, historyItem := range historyItems {
		takePillStatusMap[historyItem.PrescriptionItemID] = &Pill{
			TakeHistoryItemId: historyItem.ID,
			TakeStatus:        historyItem.TakeStatus,
			TakeDate:          r.mono.Date.Format(historyItem.TakeDate, "Y-m-d H:i:s"),
		}
	}

	// 플랜 정렬
	planKeys := make([]*Plan, 0, len(tzMap))
	for _, v := range tzMap {
		planKeys = append(planKeys, v)
	}
	sort.Slice(planKeys, func(i, j int) bool {
		org := (planKeys[i].Hour * 100) + planKeys[i].Minute
		dst := (planKeys[j].Hour * 100) + planKeys[j].Minute
		return org < dst
	})

	// 타임존을 리스트로 변환
	var plans []Plan
	for _, v := range planKeys {

		// 복용내역이 있는 경우 업데이트
		if takeStatusMap[v.TimezoneId] != nil {
			v.TakeStatus = takeStatusMap[v.TimezoneId].TakeStatus
			v.TakeDate = takeStatusMap[v.TimezoneId].TakeDate
			v.Memo = takeStatusMap[v.TimezoneId].Memo
		}

		// 의약품 복용내역이 있는 경우 업데이트
		for index, pill := range v.Pills {
			if takePillStatusMap[pill.PrescriptionItemId] != nil {
				v.Pills[index].TakeHistoryItemId = takePillStatusMap[pill.PrescriptionItemId].TakeHistoryItemId
				v.Pills[index].TakeStatus = takePillStatusMap[pill.PrescriptionItemId].TakeStatus
				v.Pills[index].TakeDate = takePillStatusMap[pill.PrescriptionItemId].TakeDate
			}
		}
		plans = append(plans, *v)
	}

	return plans, nil
}

func (r repository) GetPlansByMonth(userId uuid.UUID, year int, month int) (*Summary, error) {
	found, err := r.prescription.Query().
		Where(
			scPs.UserID(userId),
			scPs.StartedAtGTE(time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)),
			scPs.StartedAtLT(time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.UTC)),
		).
		All(r.ctx)
	if err != nil {
		return nil, err
	}

	summary := &Summary{
		Year:  fmt.Sprintf("%d", year),
		Month: fmt.Sprintf("%d", month),
	}

	for _, ps := range found {
		summaryItem := &SummaryItem{
			ID:         ps.ID,
			Name:       ps.PrescriptionName,
			StartedAt:  r.mono.Date.Format(ps.StartedAt, "Y-m-d"),
			FinishedAt: r.mono.Date.Format(ps.FinishedAt, "Y-m-d"),
			TakeDays:   ps.TakeDays,
		}
		summary.Items = append(summary.Items, *summaryItem)
	}

	return summary, nil
}

func Convert[T any, O any](data []O, action func(O) T) []T {
	var result []T
	for _, d := range data {
		result = append(result, action(d))
	}

	return result
}
