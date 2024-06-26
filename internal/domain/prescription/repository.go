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
	GetListBySearch(search *SearchCondition) ([]*Prescription, error)
	GetItemListByTimezoneLinkId(timezoneLinkId uuid.UUID) ([]*PrescriptionItem, error)
	GetItemListByTimezoneLinkIds(timezoneLinkIds []uuid.UUID) ([]*PrescriptionItem, error)
	GetItemListBySearch(search *SearchCondition) ([]*PrescriptionItem, error)
	Add(prescription *Prescription) (*Prescription, error)
	AddTx(prescription *Prescription, tx *ent.Tx) (*Prescription, error)
	Update(prescription *Prescription) (int, error)
	UpdateTx(prescription *Prescription, tx *ent.Tx) (int, error)
	Delete(id uuid.UUID) (bool, error)
	DeleteTx(id uuid.UUID, tx *ent.Tx) (bool, error)
	GetItemById(itemId uuid.UUID) (*PrescriptionItem, error)
	AddItem(item *PrescriptionItem) (*PrescriptionItem, error)
	AddItemTx(item *PrescriptionItem, tx *ent.Tx) (*PrescriptionItem, error)
	UpdateItem(prescriptionItem *PrescriptionItem) (int, error)
	DeleteItem(itemId uuid.UUID) (bool, error)
	DeleteItemByIds(itemIds []uuid.UUID) (bool, error)
	DeleteItemByIdsTx(itemIds []uuid.UUID, tx *ent.Tx) (bool, error)
}

type prescriptionRepository struct {
	root       *ent.Client
	client     *ent.PrescriptionClient
	itemClient *ent.PrescriptionItemClient
	ctx        context.Context
}

type SearchCondition struct {
	ID             uuid.UUID
	UserId         uuid.UUID
	TargetDate     time.Time
	TargetDateNext time.Time
	Limit          int
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

func (p prescriptionRepository) AddTx(prescription *Prescription, tx *ent.Tx) (*Prescription, error) {
	saved, err := tx.Prescription.
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

func (p prescriptionRepository) GetListBySearch(search *SearchCondition) ([]*Prescription, error) {
	foundList, err := p.root.Debug().Prescription.
		Query().
		Where(
			schema.And(
				schema.UserID(search.UserId),
				schema.And(
					schema.StartedAtLTE(search.TargetDate),
					schema.FinishedAtGTE(search.TargetDate),
				),
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

func (p prescriptionRepository) UpdateTx(prescription *Prescription, tx *ent.Tx) (int, error) {
	saved, err := tx.Prescription.
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
	_, err := p.client.
		Delete().
		Where(schema.ID(id)).
		Exec(p.ctx)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p prescriptionRepository) DeleteTx(id uuid.UUID, tx *ent.Tx) (bool, error) {
	_, err := tx.Prescription.
		Delete().
		Where(schema.ID(id)).
		Exec(p.ctx)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p prescriptionRepository) GetItemById(itemId uuid.UUID) (*PrescriptionItem, error) {
	found, err := p.itemClient.
		Query().
		Where(
			prescriptionitem.ID(itemId),
		).
		Only(p.ctx)
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
		SetRemainAmount(item.RemainAmount).
		SetTotalAmount(item.TotalAmount).
		SetMedicineUnit(item.MedicineUnit).
		SetMemo(item.Memo).
		SetCreatedAt(item.CreatedAt).
		Save(p.ctx)
	if err != nil {
		return nil, err
	}

	return toDomainItem(savedItem), nil
}

func (p prescriptionRepository) AddItemTx(item *PrescriptionItem, tx *ent.Tx) (*PrescriptionItem, error) {
	savedItem, err := tx.PrescriptionItem.
		Create().
		SetTimezoneLinkID(item.TimeZoneLinkId).
		SetMedicineID(item.MedicineId).
		SetMedicineName(item.MedicineName).
		SetTakeAmount(item.TakeAmount).
		SetRemainAmount(item.RemainAmount).
		SetTotalAmount(item.TotalAmount).
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
		SetTakeAmount(prescriptionItem.TakeAmount).
		SetRemainAmount(prescriptionItem.RemainAmount).
		SetMemo(prescriptionItem.Memo).
		SetUpdatedAt(time.Now().UTC()).
		Where(
			prescriptionitem.ID(prescriptionItem.ID),
		).
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

func (p prescriptionRepository) DeleteItemByIdsTx(itemIds []uuid.UUID, tx *ent.Tx) (bool, error) {
	_, err := tx.PrescriptionItem.Delete().Where(prescriptionitem.IDIn(itemIds...)).Exec(p.ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}
