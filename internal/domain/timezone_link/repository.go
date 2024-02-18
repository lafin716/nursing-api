package timezonelink

import (
	"context"
	"github.com/google/uuid"
	"nursing_api/pkg/database"
	"nursing_api/pkg/ent"
	schema "nursing_api/pkg/ent/timezonelink"
)

type repository struct {
	root         *ent.Client
	timezoneLink *ent.TimeZoneLinkClient
	c            context.Context
}

type Repository interface {
	ConnectPrescription(link *TimeZoneLink) (*TimeZoneLink, error)
	GetByPrescription(prescriptionId uuid.UUID) ([]*TimeZoneLink, error)
	DeleteByIds(ids []uuid.UUID) (bool, error)
}

func NewRepository(
	db *database.DatabaseClient,
) Repository {
	return &repository{
		root:         db.Client,
		timezoneLink: db.Client.TimeZoneLink,
		c:            db.Ctx,
	}
}

func (r repository) ConnectPrescription(link *TimeZoneLink) (*TimeZoneLink, error) {
	saved, err := r.timezoneLink.
		Create().
		SetPrescriptionID(link.PrescriptionId).
		SetTimezoneID(link.TimeZoneId).
		SetTimezoneName(link.TimeZoneName).
		SetUseAlert(link.UseAlert).
		SetMidday(link.Midday).
		SetHour(link.Hour).
		SetMinute(link.Minute).
		SetCreatedAt(link.CreatedAt).
		Save(r.c)
	if err != nil {
		return nil, err
	}

	return toDomain(saved), nil
}

func (r repository) GetByPrescription(prescriptionId uuid.UUID) ([]*TimeZoneLink, error) {
	founds, err := r.timezoneLink.
		Query().
		Where(
			schema.PrescriptionID(prescriptionId),
		).
		All(r.c)
	if err != nil {
		return nil, err
	}

	return toDomainList(founds), nil
}

func (r repository) DeleteByIds(ids []uuid.UUID) (bool, error) {
	result, err := r.timezoneLink.Delete().Where(schema.IDIn(ids...)).Exec(r.c)
	if err != nil {
		return false, err
	}

	return result > 0, nil
}
