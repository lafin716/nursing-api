package prescription

import (
	"context"
	"github.com/google/uuid"
	"log"
	"nursing_api/pkg/database"
	"nursing_api/pkg/ent"
	schema "nursing_api/pkg/ent/prescription"
	"nursing_api/pkg/ent/prescriptionitem"
	"time"
)

type prescriptionRepository struct {
	root       *ent.Client
	client     *ent.PrescriptionClient
	itemClient *ent.PrescriptionItemClient
	ctx        context.Context
}

type PrescriptionSearch struct {
	ID         uuid.UUID
	UserId     uuid.UUID
	TargetDate time.Time
	Limit      int
}

func NewPrescriptionRepository(
	dbClient *database.DatabaseClient,
) PrescriptionRepository {
	return &prescriptionRepository{
		root:       dbClient.Client,
		client:     dbClient.Client.Prescription,
		itemClient: dbClient.Client.PrescriptionItem,
		ctx:        dbClient.Ctx,
	}
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

func (p prescriptionRepository) GetListByUserId(search *PrescriptionSearch) ([]*Prescription, error) {
	foundList, err := p.client.
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

func (p prescriptionRepository) GetItemListByPrescriptionId(prescriptionId uuid.UUID) ([]*PrescriptionItem, error) {
	foundItemList, err := p.root.Debug().
		PrescriptionItem.
		Query().
		Where(
			prescriptionitem.And(
				prescriptionitem.PrescriptionID(prescriptionId),
			),
		).
		All(p.ctx)
	if err != nil {
		return nil, err
	}

	return toDomainItems(foundItemList), nil
}

func (p prescriptionRepository) Add(prescription *Prescription) (*Prescription, error) {
	tx, err := p.root.Tx(p.ctx)
	if err != nil {
		return nil, err
	}

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
		tx.Rollback()
		return nil, err
	}

	savedItems := []*ent.PrescriptionItem{}
	for _, item := range prescription.PrescriptionItems {
		saved, err := tx.PrescriptionItem.
			Create().
			SetUserID(item.UserId).
			SetMedicineID(item.MedicineId).
			SetMedicineName(item.MedicineName).
			SetTakeTimeZone(item.TakeTimeZone).
			SetTakeMoment(item.TakeMoment).
			SetTakeEtc(item.TakeEtc).
			SetTakeAmount(item.TakeAmount).
			SetMedicineUnit(item.MedicineUnit).
			SetMemo(item.Memo).
			SetCreatedAt(item.CreatedAt).
			SetPrescription(saved).
			Save(p.ctx)
		if err != nil {
			log.Printf("아이템 저장 실패:\n %+v\n", item)
			tx.Rollback()
			return nil, err
		}
		savedItems = append(savedItems, saved)
	}

	tx.Commit()
	return toDomain(saved), nil
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
		Where(prescriptionitem.PrescriptionID(id)).
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

func (p prescriptionRepository) AddItem(prescriptionId uuid.UUID, item *PrescriptionItem) (*PrescriptionItem, error) {
	savedItem, err := p.itemClient.
		Create().
		SetPrescriptionID(prescriptionId).
		SetUserID(item.UserId).
		SetMedicineID(item.MedicineId).
		SetMedicineName(item.MedicineName).
		SetTakeTimeZone(item.TakeTimeZone).
		SetTakeMoment(item.TakeMoment).
		SetTakeEtc(item.TakeEtc).
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
		SetTakeTimeZone(prescriptionItem.TakeTimeZone).
		SetTakeMoment(prescriptionItem.TakeMoment).
		SetTakeEtc(prescriptionItem.TakeEtc).
		SetTakeAmount(prescriptionItem.TakeAmount).
		SetMedicineUnit(prescriptionItem.MedicineUnit).
		SetMemo(prescriptionItem.Memo).
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
