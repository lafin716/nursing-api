package prescription

import "nursing_api/pkg/ent"

// 엔티티로 변환
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

// 엔티티 아이템으로 변환
func toEntityItems(domains []*PrescriptionItem) []*ent.PrescriptionItem {
	items := []*ent.PrescriptionItem{}
	for _, domain := range domains {
		items = append(items, toEntityItem(domain))
	}

	return items
}

func toEntityItem(domain *PrescriptionItem) *ent.PrescriptionItem {
	return &ent.PrescriptionItem{
		ID:             domain.ID,
		TimezoneLinkID: domain.TimeZoneLinkId,
		MedicineID:     domain.MedicineId,
		MedicineName:   domain.MedicineName,
		TakeAmount:     domain.TakeAmount,
		RemainAmount:   domain.RemainAmount,
		TotalAmount:    domain.TotalAmount,
		MedicineUnit:   domain.MedicineUnit,
		Memo:           domain.Memo,
		CreatedAt:      domain.CreatedAt,
	}
}

func toDomains(entities []*ent.Prescription) []*Prescription {
	domains := []*Prescription{}
	for _, entity := range entities {
		domains = append(domains, toDomain(entity))
	}

	return domains
}

func toDomain(entity *ent.Prescription) *Prescription {
	return &Prescription{
		ID:               entity.ID,
		UserId:           entity.UserID,
		PrescriptionName: entity.PrescriptionName,
		HospitalName:     entity.HospitalName,
		TakeDays:         entity.TakeDays,
		StartedAt:        entity.StartedAt,
		FinishedAt:       entity.FinishedAt,
		Memo:             entity.Memo,
		CreatedAt:        entity.CreatedAt,
		UpdatedAt:        entity.UpdatedAt,
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
		ID:             entity.ID,
		TimeZoneLinkId: entity.TimezoneLinkID,
		MedicineId:     entity.MedicineID,
		MedicineName:   entity.MedicineName,
		TakeAmount:     entity.TakeAmount,
		RemainAmount:   entity.RemainAmount,
		TotalAmount:    entity.TotalAmount,
		MedicineUnit:   entity.MedicineUnit,
		Memo:           entity.Memo,
		CreatedAt:      entity.CreatedAt,
	}
}
