package timezonelink

import (
	"context"
	"nursing_api/pkg/database"
	"nursing_api/pkg/ent"
)

type repository struct {
	root         *ent.Client
	timezoneLink *ent.TimeZoneLinkClient
	c            context.Context
}

type Repository interface {
	ConnectPrescription(link *TimeZoneLink) (*TimeZoneLink, error)
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
