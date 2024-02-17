package takehistory

import (
	"context"
	"github.com/google/uuid"
	"nursing_api/pkg/database"
	"nursing_api/pkg/ent"
	schema "nursing_api/pkg/ent/takehistory"
	"time"
)

type Repository interface {
	GetList(userId uuid.UUID) ([]*TakeHistory, error)
	GetByToday(userId uuid.UUID, today time.Time) (*TakeHistory, error)
	GetById(id uuid.UUID) (*TakeHistory, error)
	Add(newData *TakeHistory) (bool, error)
	Update(newData *TakeHistory) (bool, error)
	Delete(takeHistoryId uuid.UUID) (bool, error)
	AddItem(item *TakeHistoryItem) (bool, error)
	UpdateItem(item *TakeHistoryItem) (bool, error)
	DeleteItem(takeHistoryItemId uuid.UUID) (bool, error)
}

type repository struct {
	client     *ent.TakeHistoryClient
	itemClient *ent.TakeHistoryItemClient
	ctx        context.Context
}

func NewRepository(
	dbClient *database.DatabaseClient,
) Repository {
	return &repository{
		client:     dbClient.Client.TakeHistory,
		itemClient: dbClient.Client.TakeHistoryItem,
		ctx:        dbClient.Ctx,
	}
}
func (t repository) GetList(userId uuid.UUID) ([]*TakeHistory, error) {
	list, err := t.client.
		Query().
		Where(
			schema.UserID(userId),
		).
		All(t.ctx)
	if err != nil {
		return nil, err
	}

	return toModelList(list), nil
}

func (t repository) GetByToday(userId uuid.UUID, today time.Time) (*TakeHistory, error) {
	found, err := t.client.
		Query().
		Where(
			schema.UserID(userId),
			schema.TakeDateEQ(today),
		).
		Only(t.ctx)
	if err != nil {
		return nil, err
	}

	return toModel(found), nil
}

func (t repository) GetById(id uuid.UUID) (*TakeHistory, error) {
	found, err := t.client.
		Query().
		Where(schema.ID(id)).
		Only(t.ctx)
	if err != nil {
		return nil, err
	}

	return toModel(found), nil
}

func (t repository) Add(newData *TakeHistory) (bool, error) {
	_, err := t.client.
		Create().
		SetID(newData.ID).
		SetUserID(newData.UserId).
		SetPrescriptionID(newData.PrescriptionId).
		SetTakeDate(newData.TakeDate).
		SetTakeStatus(newData.TakeStatus).
		SetMemo(newData.Memo).
		SetCreatedAt(newData.CreatedAt).
		Save(t.ctx)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (t repository) Update(newData *TakeHistory) (bool, error) {
	result, err := t.client.
		Update().
		SetTakeDate(newData.TakeDate).
		SetTakeStatus(newData.TakeStatus).
		SetMemo(newData.Memo).
		SetUpdatedAt(newData.UpdatedAt).
		Where(
			schema.UserID(newData.ID),
		).
		Save(t.ctx)
	if err != nil {
		return false, err
	}

	return result > 0, nil
}

func (t repository) Delete(takeHistoryId uuid.UUID) (bool, error) {
	result, err := t.client.Delete().Where(schema.ID(takeHistoryId)).Exec(t.ctx)
	if err != nil {
		return false, err
	}

	return result > 0, nil
}

func (t repository) AddItem(item *TakeHistoryItem) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (t repository) UpdateItem(item *TakeHistoryItem) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (t repository) DeleteItem(takeHistoryItemId uuid.UUID) (bool, error) {
	//TODO implement me
	panic("implement me")
}
