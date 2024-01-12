package medicine

import (
	"context"
	"nursing_api/pkg/database"
	"nursing_api/pkg/ent"
	medicineSchema "nursing_api/pkg/ent/medicine"
	"time"
)

type medicineRepository struct {
	client *ent.MedicineClient
	ctx    context.Context
}

func NewAuthRepository(dbClient *database.DatabaseClient) MedicineRepository {
	return &medicineRepository{
		client: dbClient.Client.Medicine,
		ctx:    dbClient.Ctx,
	}
}

func (m medicineRepository) SavePills(medicines []*Medicine) (int, error) {
	var createRows []*ent.MedicineCreate
	for _, medicine := range medicines {
		createRows = append(createRows, m.createRow(medicine))
	}

	savedPills, err := m.client.CreateBulk(createRows...).Save(m.ctx)
	if err != nil {
		return 0, err
	}

	return len(savedPills), nil
}

func (m medicineRepository) SavePill(medicine *Medicine) (bool, error) {
	_, err := m.createRow(medicine).Save(m.ctx)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (m medicineRepository) createRow(medicine *Medicine) *ent.MedicineCreate {
	return m.client.
		Create().
		SetItemSeq(medicine.ItemSeq).
		SetMedicineName(medicine.Name).
		SetCompanyName(medicine.CompanyName).
		SetDescription(medicine.Description).
		SetUsage(medicine.Usage).
		SetEffect(medicine.Effect).
		SetSideEffect(medicine.SideEffect).
		SetCaution(medicine.Caution).
		SetWarning(medicine.Warning).
		SetInteraction(medicine.Interaction).
		SetKeepMethod(medicine.KeepMethod).
		SetAppearance(medicine.Appearance).
		SetClassName(medicine.ClassName).
		SetColorClass1(medicine.ColorClass1).
		SetColorClass2(medicine.ColorClass2).
		SetOtcName(medicine.OtcName).
		SetFormCodeName(medicine.FormCodeName).
		SetPillImage(medicine.PillImage).
		SetCreatedAt(time.Now())
}

func (m medicineRepository) GetPillsByNames(pillName string) ([]*Medicine, error) {
	foundPills := m.client.
		Query().
		Where(
			medicineSchema.MedicineNameContains(pillName),
		).
		Limit(10).
		AllX(m.ctx)

	return toDomains(foundPills), nil
}

func toDomains(entites []*ent.Medicine) []*Medicine {
	var domains []*Medicine
	for _, entity := range entites {
		domains = append(domains, toDomain(entity))
	}

	return domains
}

func toDomain(entity *ent.Medicine) *Medicine {
	return &Medicine{
		Name:         entity.MedicineName,
		ItemSeq:      entity.ItemSeq,
		CompanyName:  entity.CompanyName,
		Description:  entity.Description,
		Usage:        entity.Usage,
		Effect:       entity.Effect,
		SideEffect:   entity.SideEffect,
		Caution:      entity.Caution,
		Warning:      entity.Warning,
		Interaction:  entity.Interaction,
		KeepMethod:   entity.KeepMethod,
		Appearance:   entity.Appearance,
		ClassName:    entity.ClassName,
		ColorClass1:  entity.ColorClass1,
		ColorClass2:  entity.ColorClass2,
		PillImage:    entity.PillImage,
		OtcName:      entity.OtcName,
		FormCodeName: entity.FormCodeName,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
	}
}
