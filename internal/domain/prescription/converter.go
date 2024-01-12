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
		//PrescriptionItems: entity.PrescriptionItems,
	}
}
