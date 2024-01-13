package prescription

import "nursing_api/pkg/ent"

func toEntity(domain *Prescription) *ent.Prescription {
	return &ent.Prescription{
		ID:               domain.ID,
		UserID:           domain.UserId,
		PrescriptionName: domain.PrescriptionName,
		HospitalName:     domain.HospitalName,
		TakeDays:         domain.TakeDays,
		StartedAt:        domain.StartedAt,
		FinishedAt:       domain.FinishedAt,
		Memo:             domain.Memo,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
	}
}

func toEntityItems(domains []*PrescriptionItem) []*ent.PrescriptionItem {
	items := []*ent.PrescriptionItem{}
	for _, domain := range domains {
		items = append(items, toEntityItem(domain))
	}

	return items
}

func toEntityItem(domain *PrescriptionItem) *ent.PrescriptionItem {
	return &ent.PrescriptionItem{
		UserID:       domain.UserId,
		MedicineName: domain.MedicineName,
		TakeTimeZone: domain.TakeTimeZone,
		TakeMoment:   domain.TakeMoment,
		TakeEtc:      domain.TakeEtc,
		TakeAmount:   float64(domain.TakeAmount),
		MedicineUnit: domain.MedicineUnit,
		Memo:         domain.Memo,
		CreatedAt:    domain.CreatedAt,
	}
}

func toDomains(entities []*ent.Prescription) []*Prescription {
	domains := []*Prescription{}
	for _, entity := range entities {
		domains = append(domains, toDomain(entity))
	}

	return domains
}

// TODO 하위 연관관계 설정
func toDomain(entity *ent.Prescription) *Prescription {
	return &Prescription{
		ID:                entity.ID,
		UserId:            entity.UserID,
		PrescriptionName:  entity.PrescriptionName,
		HospitalName:      entity.HospitalName,
		TakeDays:          entity.TakeDays,
		StartedAt:         entity.StartedAt,
		FinishedAt:        entity.FinishedAt,
		Memo:              entity.Memo,
		CreatedAt:         entity.CreatedAt,
		UpdatedAt:         entity.UpdatedAt,
		PrescriptionItems: toDomainItems(entity.Edges.Items),
	}
}

func toDomainItems(entities []*ent.PrescriptionItem) []*PrescriptionItem {
	items := []*PrescriptionItem{}
	for _, entity := range entities {
		items = append(items, toDomainItem(entity))
	}

	return items
}

func toDomainItem(entity *ent.PrescriptionItem) *PrescriptionItem {
	return &PrescriptionItem{
		UserId:       entity.UserID,
		MedicineName: entity.MedicineName,
		TakeTimeZone: entity.TakeTimeZone,
		TakeMoment:   entity.TakeMoment,
		TakeEtc:      entity.TakeEtc,
		TakeAmount:   entity.TakeAmount,
		MedicineUnit: entity.MedicineUnit,
		Memo:         entity.Memo,
		CreatedAt:    entity.CreatedAt,
	}
}
