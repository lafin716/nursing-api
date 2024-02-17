package timezone

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"nursing_api/pkg/database"
	"nursing_api/pkg/ent"
	schemaTimezone "nursing_api/pkg/ent/timezone"
	"time"
)

type Repository interface {
	GetTimeZones(userId uuid.UUID) ([]*TimeZone, error)
	GetTimeZone(id uuid.UUID, userId uuid.UUID) (*TimeZone, error)
	CreateTimeZone(model *TimeZone) (*TimeZone, error)
	UpdateTimeZone(model *TimeZone) (bool, error)
	DeleteTimeZone(id uuid.UUID, userId uuid.UUID) (bool, error)
	GetDuplicate(userId uuid.UUID, name string, midday string, hour string, minute string) (*TimeZone, error)
}

type timezoneRepository struct {
	root     *ent.Client
	timezone *ent.TimeZoneClient
	c        context.Context
}

func NewRepository(
	dbClient *database.DatabaseClient,
) Repository {
	return &timezoneRepository{
		root:     dbClient.Client,
		timezone: dbClient.Client.TimeZone,
		c:        dbClient.Ctx,
	}
}

func (p timezoneRepository) GetTimeZones(userId uuid.UUID) ([]*TimeZone, error) {
	list, err := p.timezone.
		Query().
		Where(
			schemaTimezone.UserID(userId),
		).
		Order(
			schemaTimezone.ByMidday(sql.OrderAsc()),
			schemaTimezone.ByHour(sql.OrderAsc()),
			schemaTimezone.ByMinute(sql.OrderAsc()),
		).
		All(p.c)
	if err != nil {
		return nil, err
	}

	return toTimeZoneDomainList(list), nil
}

func (p timezoneRepository) GetTimeZone(id uuid.UUID, userId uuid.UUID) (*TimeZone, error) {
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

func (p timezoneRepository) CreateTimeZone(model *TimeZone) (*TimeZone, error) {
	newPlanTimeZone, err := p.timezone.
		Create().
		SetUserID(model.UserID).
		SetTimezoneName(model.Name).
		SetIsDefault(model.IsDefault).
		SetMidday(model.Midday).
		SetHour(model.Hour).
		SetMinute(model.Minute).
		SetCreatedAt(time.Now()).
		Save(p.c)
	if err != nil {
		return nil, err
	}

	return toTimeZoneDomain(newPlanTimeZone), nil
}

func (p timezoneRepository) UpdateTimeZone(model *TimeZone) (bool, error) {
	err := p.timezone.
		Update().
		SetTimezoneName(model.Name).
		SetMidday(model.Midday).
		SetHour(model.Hour).
		SetMinute(model.Minute).
		Where(schemaTimezone.ID(model.ID)).
		Exec(p.c)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p timezoneRepository) DeleteTimeZone(id uuid.UUID, userId uuid.UUID) (bool, error) {
	deleted, err := p.timezone.
		Delete().
		Where(schemaTimezone.ID(id), schemaTimezone.UserID(userId)).
		Exec(p.c)
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
	found, err := p.root.Debug().TimeZone.
		Query().
		Where(
			schemaTimezone.And(
				schemaTimezone.UserID(userId),
				schemaTimezone.Or(
					schemaTimezone.TimezoneNameEQ(name),
					schemaTimezone.And(
						schemaTimezone.Midday(midday),
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
