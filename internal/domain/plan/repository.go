package plan

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"nursing_api/pkg/database"
	"nursing_api/pkg/ent"
	scPs "nursing_api/pkg/ent/prescription"
	scPsItem "nursing_api/pkg/ent/prescriptionitem"
	scTzl "nursing_api/pkg/ent/timezonelink"
	"time"
)

type Repository interface {
	GetPlans(userId uuid.UUID, date time.Time) ([]Plan, error)
}

type repository struct {
	root             *ent.Client
	prescription     *ent.PrescriptionClient
	prescriptionItem *ent.PrescriptionItemClient
	timezoneLink     *ent.TimeZoneLinkClient
	ctx              context.Context
}

func NewRepository(
	db *database.DatabaseClient,
) Repository {
	return &repository{
		root:             db.Client,
		prescription:     db.Client.Prescription,
		prescriptionItem: db.Client.PrescriptionItem,
		timezoneLink:     db.Client.TimeZoneLink,
		ctx:              db.Ctx,
	}
}

func (r repository) GetPlans(userId uuid.UUID, date time.Time) ([]Plan, error) {
	// 현재날짜에 유효한 처방전 조회
	pcs, err := r.root.Debug().Prescription.Query().Where(
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

	pcsIds := Convert[uuid.UUID, *ent.Prescription](pcs, func(data *ent.Prescription) uuid.UUID { return data.ID })

	// 처방전과 연결된 타임존 조회
	tzls, err := r.root.Debug().TimeZoneLink.
		Query().
		Where(
			scTzl.PrescriptionIDIn(pcsIds...),
		).All(r.ctx)

	// 타임존 고유번호로 타임존링크 고유번호 그룹핑
	tzMap := map[uuid.UUID]*Plan{}
	for _, tzl := range tzls {
		pcsItems, err := r.root.Debug().PrescriptionItem.Query().Where(scPsItem.TimezoneLinkID(tzl.ID)).All(r.ctx)
		if err != nil {
			return []Plan{}, err
		}

		var pills []Pill
		for _, pcsItem := range pcsItems {
			pills = append(pills, Pill{
				PillName:   pcsItem.MedicineName,
				MedicineId: pcsItem.MedicineID,
				TakeUnit:   pcsItem.MedicineUnit,
				TakeAmount: pcsItem.TakeAmount,
			})
		}

		var plan *Plan
		if tzMap[tzl.TimezoneID] == nil {
			tzText := fmt.Sprintf("%s %s:%s", tzl.Midday, tzl.Hour, tzl.Minute)
			plan = &Plan{
				TimezoneId: tzl.TimezoneID,
				PlanName:   tzl.TimezoneName,
				Timezone:   tzText,
			}
		} else {
			plan = tzMap[tzl.TimezoneID]
		}

		plan.Pills = append(plan.Pills, pills...)
		tzMap[tzl.TimezoneID] = plan
	}

	// 타임존을 리스트로 변환
	var plans []Plan
	for _, v := range tzMap {
		plans = append(plans, *v)
	}

	return plans, nil
}

func Convert[T any, O any](data []O, action func(O) T) []T {
	var result []T
	for _, d := range data {
		result = append(result, action(d))
	}

	return result
}
