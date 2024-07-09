package prescription

import (
	"context"
	"github.com/google/uuid"
	"nursing_api/pkg/database"
	"nursing_api/pkg/ent"
	pscSchema "nursing_api/pkg/ent/prescription"
	pscItmSchema "nursing_api/pkg/ent/prescriptionitem"
	"time"
)

type Repository interface {
	// 처방전
	GetList(search *SearchCondition) ([]*Prescription, error)
	GetById(id uuid.UUID) (*Prescription, error)
	Add(prescription *Prescription) (*Prescription, error)
	Update(prescription *Prescription) (int, error)
	Delete(id uuid.UUID) (bool, error)

	// 처방전 아이템
	GetItemList(prescriptionId uuid.UUID) ([]*PrescriptionItem, error)
	GetItemById(itemId uuid.UUID) (*PrescriptionItem, error)
	AddItem(item *PrescriptionItem) (*PrescriptionItem, error)
	UpdateItem(prescriptionItem *PrescriptionItem) (int, error)
	DeleteItem(itemId uuid.UUID) (bool, error)
}

type prescriptionRepository struct {
	db database.DatabaseClient
}

type SearchCondition struct {
	ID             uuid.UUID
	UserId         uuid.UUID
	TargetDate     time.Time
	TargetDateNext time.Time
	Limit          int
}

func NewRepository(
	db database.DatabaseClient,
) Repository {
	return &prescriptionRepository{
		db: db,
	}
}

func (p prescriptionRepository) GetTxManager() database.TransactionManager {
	return p.db.GetTxManager()
}

func (p prescriptionRepository) GetPrescriptionClient() *ent.PrescriptionClient {
	return p.db.GetClient().Prescription
}

func (p prescriptionRepository) GetPrescriptionItemClient() *ent.PrescriptionItemClient {
	return p.db.GetClient().PrescriptionItem
}

func (p prescriptionRepository) GetCtx() context.Context {
	return p.db.GetCtx()
}

func (p prescriptionRepository) GetList(search *SearchCondition) ([]*Prescription, error) {
	foundList, err := p.GetPrescriptionClient().
		Query().
		Where(
			pscSchema.And(
				pscSchema.UserID(search.UserId),
				pscSchema.And(
					pscSchema.StartedAtLTE(search.TargetDate),
					pscSchema.FinishedAtGTE(search.TargetDate),
				),
			),
		).
		Limit(search.Limit).
		All(p.GetCtx())
	if err != nil {
		return nil, err
	}

	return toDomains(foundList), nil
}

func (p prescriptionRepository) GetById(id uuid.UUID) (*Prescription, error) {
	found, err := p.GetPrescriptionClient().
		Query().
		Where(
			pscSchema.ID(id),
		).
		Only(p.GetCtx())
	if err != nil {
		return nil, err
	}

	return toDomain(found), nil
}

func (p prescriptionRepository) Add(prescription *Prescription) (*Prescription, error) {
	saved, err := p.GetPrescriptionClient().
		Create().
		SetUserID(prescription.UserId).
		SetPrescriptionName(prescription.PrescriptionName).
		SetHospitalName(prescription.HospitalName).
		SetTakeDays(prescription.TakeDays).
		SetStartedAt(prescription.StartedAt).
		SetFinishedAt(prescription.FinishedAt).
		SetCreatedAt(prescription.CreatedAt).
		Save(p.GetCtx())
	if err != nil {
		return nil, err
	}

	return toDomain(saved), nil
}

func (p prescriptionRepository) Update(prescription *Prescription) (int, error) {
	saved, err := p.GetPrescriptionClient().
		Update().
		SetPrescriptionName(prescription.PrescriptionName).
		SetHospitalName(prescription.HospitalName).
		SetTakeDays(prescription.TakeDays).
		SetStartedAt(prescription.StartedAt).
		SetFinishedAt(prescription.FinishedAt).
		SetUpdatedAt(time.Now()).
		Where(
			pscSchema.ID(prescription.ID),
		).
		Save(p.GetCtx())
	if err != nil {
		return 0, err
	}

	return saved, nil
}

func (p prescriptionRepository) Delete(id uuid.UUID) (bool, error) {
	_, err := p.GetPrescriptionClient().
		Delete().
		Where(pscSchema.ID(id)).
		Exec(p.GetCtx())
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p prescriptionRepository) GetItemList(prescriptionId uuid.UUID) ([]*PrescriptionItem, error) {
	foundList, err := p.GetPrescriptionItemClient().
		Query().
		Where(
			pscItmSchema.PrescriptionIDEQ(prescriptionId),
		).
		All(p.GetCtx())
	if err != nil {
		return nil, err
	}

	return toDomainItems(foundList), nil
}

func (p prescriptionRepository) GetItemById(itemId uuid.UUID) (*PrescriptionItem, error) {
	found, err := p.GetPrescriptionItemClient().
		Query().
		Where(
			pscItmSchema.ID(itemId),
		).
		Only(p.GetCtx())
	if err != nil {
		return nil, err
	}

	return toDomainItem(found), nil
}

func (p prescriptionRepository) AddItem(item *PrescriptionItem) (*PrescriptionItem, error) {
	savedItem, err := p.GetPrescriptionItemClient().
		Create().
		SetPrescriptionID(item.PrescriptionId).
		SetTimezoneID(item.TimeZoneId).
		SetTimezoneName(item.TimeZoneName).
		SetMidday(item.Midday).
		SetHour(item.Hour).
		SetMinute(item.Minute).
		SetMedicineID(item.MedicineId).
		SetMedicineName(item.MedicineName).
		SetTakeAmount(item.TakeAmount).
		SetTotalAmount(item.TotalAmount).
		SetMedicineUnit(item.MedicineUnit).
		SetMemo(item.Memo).
		SetCreatedAt(item.CreatedAt).
		Save(p.GetCtx())
	if err != nil {
		return nil, err
	}

	return toDomainItem(savedItem), nil
}

func (p prescriptionRepository) UpdateItem(prescriptionItem *PrescriptionItem) (int, error) {
	err := p.GetPrescriptionItemClient().
		Update().
		SetTotalAmount(prescriptionItem.TotalAmount).
		SetRemainAmount(prescriptionItem.RemainAmount).
		SetTakeAmount(prescriptionItem.TakeAmount).
		SetMemo(prescriptionItem.Memo).
		SetUpdatedAt(time.Now()).
		Where(
			pscItmSchema.ID(prescriptionItem.ID),
		).
		Exec(p.GetCtx())
	if err != nil {
		return 0, err
	}

	return 1, nil
}

func (p prescriptionRepository) DeleteItem(itemId uuid.UUID) (bool, error) {
	_, err := p.GetPrescriptionItemClient().
		Delete().
		Where(pscItmSchema.ID(itemId)).
		Exec(p.GetCtx())
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p prescriptionRepository) DeleteItemByIds(itemIds []uuid.UUID) (bool, error) {
	_, err := p.GetPrescriptionItemClient().
		Delete().
		Where(pscItmSchema.IDIn(itemIds...)).
		Exec(p.GetCtx())
	if err != nil {
		return false, err
	}
	return true, nil
}
