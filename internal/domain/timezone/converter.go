package timezone

import "nursing_api/pkg/ent"

func toTimeZoneDomainList(entities []*ent.TimeZone) []*TimeZone {
	domains := []*TimeZone{}
	for _, entity := range entities {
		domains = append(domains, toTimeZoneDomain(entity))
	}

	return domains
}

func toTimeZoneDomain(entity *ent.TimeZone) *TimeZone {
	return &TimeZone{
		ID:        entity.ID,
		UserID:    entity.UserID,
		Name:      entity.TimezoneName,
		IsDefault: entity.IsDefault,
		Midday:    entity.Midday,
		Hour:      entity.Hour,
		Minute:    entity.Minute,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
