package plan

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"nursing_api/pkg/database"
	"nursing_api/pkg/ent"
	schemaTimezone "nursing_api/pkg/ent/plantimezone"
	"time"
)

type planRepository struct {
	root     *ent.Client
	timezone *ent.PlanTimeZoneClient
	c        context.Context
}

func NewPlanRepository(
	dbClient *database.DatabaseClient,
) PlanRepository {
	return &planRepository{
		root:     dbClient.Client,
		timezone: dbClient.Client.PlanTimeZone,
		c:        dbClient.Ctx,
	}
}

func (p planRepository) GetTimeZones(userId uuid.UUID) ([]*TimeZone, error) {
	list, err := p.timezone.
		Query().
		Where(
			schemaTimezone.UserID(userId),
		).
		Order(
			schemaTimezone.ByMeridiem(sql.OrderAsc()),
			schemaTimezone.ByHour(sql.OrderAsc()),
			schemaTimezone.ByMinute(sql.OrderAsc()),
		).
		All(p.c)
	if err != nil {
		return nil, err
	}

	return toTimeZoneDomainList(list), nil
}

func (p planRepository) GetTimeZone(id uuid.UUID, userId uuid.UUID) (*TimeZone, error) {
	timezone, err := p.timezone.
		Query().
		Where(
			schemaTimezone.ID(id),
			schemaTimezone.UserID(userId),
		).
		Only(p.c)
	if err != nil {
		return nil, err
	}

	return toTimeZoneDomain(timezone), nil
}

func (p planRepository) CreateTimeZone(model *TimeZone) (*TimeZone, error) {
	newPlanTimeZone, err := p.timezone.
		Create().
		SetUserID(model.UserID).
		SetTimezoneName(model.Name).
		SetIsDefault(model.IsDefault).
		SetUseAlert(model.UseAlert).
		SetMeridiem(model.Meridiem).
		SetHour(model.Hour).
		SetMinute(model.Minute).
		SetCreatedAt(time.Now()).
		Save(p.c)
	if err != nil {
		return nil, err
	}

	return toTimeZoneDomain(newPlanTimeZone), nil
}

func (p planRepository) UpdateTimeZone(model *TimeZone) (bool, error) {
	err := p.timezone.
		Update().
		SetTimezoneName(model.Name).
		SetUseAlert(model.UseAlert).
		SetMeridiem(model.Meridiem).
		SetHour(model.Hour).
		SetMinute(model.Minute).
		Where(schemaTimezone.ID(model.ID)).
		Exec(p.c)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p planRepository) DeleteTimeZone(id uuid.UUID, userId uuid.UUID) (bool, error) {
	deleted, err := p.timezone.
		Delete().
		Where(schemaTimezone.ID(id), schemaTimezone.UserID(userId)).
		Exec(p.c)
	if err != nil {
		return false, err
	}

	return deleted > 0, nil
}

func (p planRepository) GetDuplicate(
	userId uuid.UUID,
	name string,
	meridiem string,
	hour string,
	minute string,
) (*TimeZone, error) {
	found, err := p.root.Debug().PlanTimeZone.
		Query().
		Where(
			schemaTimezone.And(
				schemaTimezone.UserID(userId),
				schemaTimezone.Or(
					schemaTimezone.TimezoneNameEQ(name),
					schemaTimezone.And(
						schemaTimezone.Meridiem(meridiem),
						schemaTimezone.Hour(hour),
						schemaTimezone.Minute(minute),
					),
				),
			),
		).
		Only(p.c)
	if err != nil {
		return nil, err
	}

	return toTimeZoneDomain(found), nil
}
