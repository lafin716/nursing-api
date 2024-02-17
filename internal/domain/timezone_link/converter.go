package timezonelink

import "nursing_api/pkg/ent"

func toDomainList(entities []*ent.TimeZoneLink) []*TimeZoneLink {
	domains := []*TimeZoneLink{}
	for _, entity := range entities {
		domains = append(domains, toDomain(entity))
	}

	return domains
}

func toDomain(entity *ent.TimeZoneLink) *TimeZoneLink {
	return &TimeZoneLink{
		ID:             entity.ID,
		PrescriptionId: entity.PrescriptionID,
		TimeZoneId:     entity.TimezoneID,
		TimeZoneName:   entity.TimezoneName,
		UseAlert:       entity.UseAlert,
		Midday:         entity.Midday,
		Hour:           entity.Hour,
		Minute:         entity.Minute,
		CreatedAt:      entity.CreatedAt,
		UpdatedAt:      entity.UpdatedAt,
	}
}
