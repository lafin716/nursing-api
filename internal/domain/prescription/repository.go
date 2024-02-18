package prescription

import (
	"context"
	"github.com/google/uuid"
	"nursing_api/pkg/database"
	"nursing_api/pkg/ent"
	schema "nursing_api/pkg/ent/prescription"
	"nursing_api/pkg/ent/prescriptionitem"
	"time"
)

type Repository interface {
	GetById(id uuid.UUID) (*Prescription, error)
	GetListByUserId(search *SearchCondition) ([]*Prescription, error)
	GetItemListByTimezoneLinkId(timezoneLinkId uuid.UUID) ([]*PrescriptionItem, error)
	GetItemListByTimezoneLinkIds(timezoneLinkIds []uuid.UUID) ([]*PrescriptionItem, error)
	GetItemListBySearch(search *SearchCondition) ([]*PrescriptionItem, error)
	Add(prescription *Prescription) (*Prescription, error)
	Update(prescription *Prescription) (int, error)
	Delete(id uuid.UUID) (bool, error)
	GetItemById(itemId uuid.UUID) (*PrescriptionItem, error)
	AddItem(item *PrescriptionItem) (*PrescriptionItem, error)
	UpdateItem(prescriptionItem *PrescriptionItem) (int, error)
	DeleteItem(itemId uuid.UUID) (bool, error)
	DeleteItemByIds(itemIds []uuid.UUID) (bool, error)
}

type prescriptionRepository struct {
	root       *ent.Client
	client     *ent.PrescriptionClient
	itemClient *ent.PrescriptionItemClient
	ctx        context.Context
}

type SearchCondition struct {
	ID         uuid.UUID
	UserId     uuid.UUID
	TargetDate time.Time
	Limit      int
}

func NewRepository(
	dbClient *database.DatabaseClient,
) Repository {
	return &prescriptionRepository{
		root:       dbClient.Client,
		client:     dbClient.Client.Prescription,
		itemClient: dbClient.Client.PrescriptionItem,
		ctx:        dbClient.Ctx,
	}
}

func (p prescriptionRepository) Add(prescription *Prescription) (*Prescription, error) {
	saved, err := p.client.
		Create().
		SetUserID(prescription.UserId).
		SetPrescriptionName(prescription.PrescriptionName).
		SetHospitalName(prescription.HospitalName).
		SetTakeDays(prescription.TakeDays).
		SetStartedAt(prescription.StartedAt).
		SetFinishedAt(prescription.FinishedAt).
		SetMemo(prescription.Memo).
		SetCreatedAt(prescription.CreatedAt).
		Save(p.ctx)
	if err != nil {
		return nil, err
	}

	return toDomain(saved), nil
}

func (p prescriptionRepository) GetById(id uuid.UUID) (*Prescription, error) {
	found, err := p.client.
		Query().
		Where(
			schema.ID(id),
		).
		Only(p.ctx)
	if err != nil {
		return nil, err
	}

	return toDomain(found), nil
}

func (p prescriptionRepository) GetListByUserId(search *SearchCondition) ([]*Prescription, error) {
	foundList, err := p.root.Debug().Prescription.
		Query().
		Where(
			schema.And(
				schema.UserID(search.UserId),
				schema.StartedAtLTE(search.TargetDate),
				schema.FinishedAtGTE(search.TargetDate),
			),
		).
		Limit(search.Limit).
		All(p.ctx)
	if err != nil {
		return nil, err
	}

	return toDomains(foundList), nil
}

func (p prescriptionRepository) GetItemListByTimezoneLinkId(timezoneLinkId uuid.UUID) ([]*PrescriptionItem, error) {
	foundItemList, err := p.root.Debug().
		PrescriptionItem.
		Query().
		Where(
			prescriptionitem.TimezoneLinkID(timezoneLinkId),
		).
		All(p.ctx)
	if err != nil {
		return nil, err
	}

	return toDomainItems(foundItemList), nil
}

func (p prescriptionRepository) GetItemListByTimezoneLinkIds(timezoneLinkIds []uuid.UUID) ([]*PrescriptionItem, error) {
	foundItemList, err := p.root.Debug().
		PrescriptionItem.
		Query().
		Where(
			prescriptionitem.TimezoneLinkIDIn(timezoneLinkIds...),
		).
		All(p.ctx)
	if err != nil {
		return nil, err
	}

	return toDomainItems(foundItemList), nil
}

func (p prescriptionRepository) GetItemListBySearch(search *SearchCondition) ([]*PrescriptionItem, error) {
	found, err := p.root.Debug().PrescriptionItem.
		Query().
		All(p.ctx)
	if err != nil {
		return nil, err
	}

	return toDomainItems(found), nil
}

func (p prescriptionRepository) Update(prescription *Prescription) (int, error) {
	saved, err := p.client.
		Update().
		SetPrescriptionName(prescription.PrescriptionName).
		SetHospitalName(prescription.HospitalName).
		SetTakeDays(prescription.TakeDays).
		SetStartedAt(prescription.StartedAt).
		SetFinishedAt(prescription.FinishedAt).
		SetMemo(prescription.Memo).
		SetUpdatedAt(time.Now()).
		Where(
			schema.ID(prescription.ID),
		).
		Save(p.ctx)
	if err != nil {
		return 0, err
	}

	return saved, nil
}

func (p prescriptionRepository) Delete(id uuid.UUID) (bool, error) {
	tx, err := p.root.Tx(p.ctx)
	if err != nil {
		return false, err
	}

	// 하위 아이템 모두 삭제
	_, err = tx.PrescriptionItem.
		Delete().
		Exec(p.ctx)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	_, err = tx.Prescription.
		Delete().
		Where(schema.ID(id)).
		Exec(p.ctx)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()
	return true, nil
}

func (p prescriptionRepository) GetItemById(itemId uuid.UUID) (*PrescriptionItem, error) {
	found, err := p.itemClient.Get(p.ctx, itemId)
	if err != nil {
		return nil, err
	}

	return toDomainItem(found), nil
}

func (p prescriptionRepository) AddItem(item *PrescriptionItem) (*PrescriptionItem, error) {
	savedItem, err := p.itemClient.
		Create().
		SetTimezoneLinkID(item.TimeZoneLinkId).
		SetMedicineID(item.MedicineId).
		SetMedicineName(item.MedicineName).
		SetTakeAmount(item.TakeAmount).
		SetMedicineUnit(item.MedicineUnit).
		SetMemo(item.Memo).
		SetCreatedAt(item.CreatedAt).
		Save(p.ctx)
	if err != nil {
		return nil, err
	}

	return toDomainItem(savedItem), nil
}

func (p prescriptionRepository) UpdateItem(prescriptionItem *PrescriptionItem) (int, error) {
	err := p.itemClient.
		Update().
		SetMedicineID(prescriptionItem.MedicineId).
		SetMedicineName(prescriptionItem.MedicineName).
		Exec(p.ctx)
	if err != nil {
		return 0, err
	}

	return 1, nil
}

func (p prescriptionRepository) DeleteItem(itemId uuid.UUID) (bool, error) {
	_, err := p.itemClient.Delete().Where(prescriptionitem.ID(itemId)).Exec(p.ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p prescriptionRepository) DeleteItemByIds(itemIds []uuid.UUID) (bool, error) {
	_, err := p.itemClient.Delete().Where(prescriptionitem.IDIn(itemIds...)).Exec(p.ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}
