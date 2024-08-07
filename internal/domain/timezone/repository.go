package timezone

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"nursing_api/pkg/database"
	"nursing_api/pkg/ent"
	pscSchema "nursing_api/pkg/ent/prescription"
	pscItmSchema "nursing_api/pkg/ent/prescriptionitem"
	tkhItmSchema "nursing_api/pkg/ent/takehistoryitem"
	tzSchema "nursing_api/pkg/ent/timezone"
	"nursing_api/pkg/mono"
	"time"
)

type Repository interface {
	GetTxManager() database.TransactionManager

	GetTimeZones(userId uuid.UUID) ([]*TimeZone, error)
	GetTimeZone(id uuid.UUID, userId uuid.UUID) (*TimeZone, error)
	CreateTimeZone(model *TimeZone) (*TimeZone, error)
	UpdateTimeZone(model *TimeZone) (bool, error)
	DeleteTimeZone(id uuid.UUID, userId uuid.UUID) (bool, error)
	GetDuplicate(userId uuid.UUID, name string, midday string, hour string, minute string) (*TimeZone, error)
	CountPrescriptionItemByTimeZoneId(userId uuid.UUID, timezoneId uuid.UUID, date time.Time) (int, error)
	UpdateCurrentPrescriptionItemByTimeZone(userId uuid.UUID, model *TimeZone) error
	UpdateCurrentTakeHistoryItemByTimeZone(userId uuid.UUID, model *TimeZone) error
}

type timezoneRepository struct {
	db   database.DatabaseClient
	mono *mono.Client
}

func NewRepository(
	db database.DatabaseClient,
	mono *mono.Client,
) Repository {
	return &timezoneRepository{
		db:   db,
		mono: mono,
	}
}

func (p timezoneRepository) GetTimeZones(userId uuid.UUID) ([]*TimeZone, error) {
	list, err := p.timezoneClient().
		Query().
		Where(
			tzSchema.UserID(userId),
		).
		Order(
			tzSchema.ByMidday(sql.OrderAsc()),
			tzSchema.ByHour(sql.OrderAsc()),
			tzSchema.ByMinute(sql.OrderAsc()),
		).
		All(p.GetCtx())
	if err != nil {
		return nil, err
	}

	return toTimeZoneDomainList(list), nil
}

func (p timezoneRepository) GetTimeZone(id uuid.UUID, userId uuid.UUID) (*TimeZone, error) {
	timezone, err := p.timezoneClient().
		Query().
		Where(
			tzSchema.ID(id),
			tzSchema.UserID(userId),
		).
		Only(p.GetCtx())
	if err != nil {
		return nil, err
	}

	return toTimeZoneDomain(timezone), nil
}

func (p timezoneRepository) CreateTimeZone(model *TimeZone) (*TimeZone, error) {
	newPlanTimeZone, err := p.timezoneClient().
		Create().
		SetUserID(model.UserID).
		SetTimezoneName(model.Name).
		SetIsDefault(model.IsDefault).
		SetMidday(model.Midday).
		SetHour(model.Hour).
		SetMinute(model.Minute).
		SetCreatedAt(time.Now()).
		Save(p.GetCtx())
	if err != nil {
		return nil, err
	}

	return toTimeZoneDomain(newPlanTimeZone), nil
}

func (p timezoneRepository) UpdateTimeZone(model *TimeZone) (bool, error) {
	err := p.timezoneClient().
		Update().
		SetTimezoneName(model.Name).
		SetMidday(model.Midday).
		SetHour(model.Hour).
		SetMinute(model.Minute).
		Where(tzSchema.ID(model.ID)).
		Exec(p.GetCtx())
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p timezoneRepository) DeleteTimeZone(id uuid.UUID, userId uuid.UUID) (bool, error) {
	deleted, err := p.timezoneClient().
		Delete().
		Where(tzSchema.ID(id), tzSchema.UserID(userId)).
		Exec(p.GetCtx())
	if err != nil {
		return false, err
	}

	return deleted > 0, nil
}

func (p timezoneRepository) GetDuplicate(
	userId uuid.UUID,
	name string,
	midday string,
	hour string,
	minute string,
) (*TimeZone, error) {
	found, err := p.timezoneClient().
		Query().
		Where(
			tzSchema.And(
				tzSchema.UserID(userId),
				tzSchema.Or(
					tzSchema.TimezoneNameEQ(name),
					tzSchema.And(
						tzSchema.Midday(midday),
						tzSchema.Hour(hour),
						tzSchema.Minute(minute),
					),
				),
			),
		).
		Only(p.GetCtx())
	if err != nil {
		return nil, err
	}

	return toTimeZoneDomain(found), nil
}

func (p timezoneRepository) CountPrescriptionItemByTimeZoneId(
	userId uuid.UUID,
	timezoneId uuid.UUID,
	date time.Time,
) (int, error) {
	psc, err := p.prescriptionClient().
		Query().
		WithPrescriptionItems(func(query *ent.PrescriptionItemQuery) {
			query.Where(
				pscItmSchema.TimezoneIDEQ(timezoneId),
			)
		}).
		Where(
			pscSchema.And(
				pscSchema.UserID(userId),
				pscSchema.StartedAtLTE(p.mono.Date.TruncateToDate(date)),
				pscSchema.FinishedAtGT(p.mono.Date.TruncateToDateAddDay(date, 1)),
			),
		).
		All(p.GetCtx())
	if err != nil {
		return 0, err
	}

	count := 0
	for _, item := range psc {
		count += len(item.Edges.PrescriptionItems)
	}

	return count, nil
}

// 금일자 의약품아이템과 내역을 업데이트
func (p timezoneRepository) UpdateCurrentPrescriptionItemByTimeZone(userId uuid.UUID, timezone *TimeZone) error {
	updated := p.prescriptionItemClient().
		Update().
		SetTimezoneName(timezone.Name).
		SetMidday(timezone.Midday).
		SetHour(timezone.Hour).
		SetMinute(timezone.Minute).
		Where(
			pscItmSchema.UserIDEQ(userId),
			pscItmSchema.TimezoneIDEQ(timezone.ID),
		).Exec(p.GetCtx())

	return updated
}

// 금일자 복용내역을 업데이트
func (p timezoneRepository) UpdateCurrentTakeHistoryItemByTimeZone(userId uuid.UUID, timezone *TimeZone) error {
	updated := p.takeHistoryItemClient().
		Update().
		SetTimezoneName(timezone.Name).
		SetMidday(timezone.Midday).
		SetHour(timezone.Hour).
		SetMinute(timezone.Minute).
		Where(
			tkhItmSchema.And(
				tkhItmSchema.UserIDEQ(userId),
				tkhItmSchema.TimezoneIDEQ(timezone.ID),
				tkhItmSchema.TakeDateEQ(p.mono.Date.TruncateToDate(time.Now()).Format("2006-01-02")),
			),
		).Exec(p.GetCtx())

	return updated
}

func (p timezoneRepository) timezoneClient() *ent.TimeZoneClient {
	return p.db.GetClient().TimeZone
}

func (p timezoneRepository) prescriptionClient() *ent.PrescriptionClient {
	return p.db.GetClient().Debug().Prescription
}

func (p timezoneRepository) prescriptionItemClient() *ent.PrescriptionItemClient {
	return p.db.GetClient().PrescriptionItem
}

func (p timezoneRepository) takeHistoryItemClient() *ent.TakeHistoryItemClient {
	return p.db.GetClient().TakeHistoryItem
}

func (p timezoneRepository) GetCtx() context.Context {
	return p.db.GetCtx()
}

func (p timezoneRepository) GetTxManager() database.TransactionManager {
	return p.db.GetTxManager()
}
