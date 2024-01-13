package prescription

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"nursing_api/pkg/database"
	"nursing_api/pkg/ent"
	schema "nursing_api/pkg/ent/prescription"
	"time"
)

type prescriptionRepository struct {
	client *ent.PrescriptionClient
	ctx    context.Context
}

func NewPrescriptionRepository(
	dbClient *database.DatabaseClient,
) PrescriptionRepository {
	return &prescriptionRepository{
		client: dbClient.Client.Prescription,
		ctx:    dbClient.Ctx,
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

func (p prescriptionRepository) GetListByUserId(userId uuid.UUID) ([]*Prescription, error) {
	foundList, err := p.client.
		Query().
		Where(schema.UserID(userId)).
		Limit(10).
		All(p.ctx)
	if err != nil {
		return nil, err
	}

	return toDomains(foundList), nil
}

// TODO 연관관계 설정 후 구현
func (p prescriptionRepository) GetItemListByPrescriptionId(prescriptionId uuid.UUID) ([]*Prescription, error) {
	return nil, errors.New("아직 미구현")
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

func (p prescriptionRepository) AddItem(item PrescriptionItem) (*PrescriptionItem, error) {
	return nil, errors.New("아직 미구현")
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

func (p prescriptionRepository) UpdateItem(itemId uuid.UUID) (int, error) {
	return 0, errors.New("아직 미구현")
}

func (p prescriptionRepository) Delete(id uuid.UUID) (bool, error) {
	_, err := p.client.Delete().Where(schema.ID(id)).Exec(p.ctx)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p prescriptionRepository) DeleteItem(itemId uuid.UUID) (bool, error) {
	return false, errors.New("아직 미구현")
}
