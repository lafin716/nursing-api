package plan

import "nursing_api/pkg/ent"

func toTimeZoneDomainList(entities []*ent.PlanTimeZone) []*PlanTimeZone {
	domains := []*PlanTimeZone{}
	for _, entity := range entities {
		domains = append(domains, toTimeZoneDomain(entity))
	}

	return domains
}

func toTimeZoneDomain(entity *ent.PlanTimeZone) *PlanTimeZone {
	return &PlanTimeZone{
		ID:        entity.ID,
		UserID:    entity.UserID,
		Name:      entity.TimezoneName,
		IsDefault: entity.IsDefault,
		UseAlert:  entity.UseAlert,
		Meridiem:  entity.Meridiem,
		Hour:      entity.Hour,
		Minute:    entity.Minute,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
