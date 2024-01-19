package take_history

import (
	"context"
	"github.com/google/uuid"
	"nursing_api/pkg/database"
	"nursing_api/pkg/ent"
	schema "nursing_api/pkg/ent/takehistory"
	"time"
)

type takeHistoryRepository struct {
	client     *ent.TakeHistoryClient
	itemClient *ent.TakeHistoryItemClient
	ctx        context.Context
}

func NewTakeHistoryRepository(
	dbClient *database.DatabaseClient,
) TakeHistoryRepository {
	return &takeHistoryRepository{
		client:     dbClient.Client.TakeHistory,
		itemClient: dbClient.Client.TakeHistoryItem,
		ctx:        dbClient.Ctx,
	}
}
func (t takeHistoryRepository) GetList(userId uuid.UUID) ([]*TakeHistory, error) {
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

func (t takeHistoryRepository) GetByToday(userId uuid.UUID, today time.Time) (*TakeHistory, error) {
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

func (t takeHistoryRepository) GetById(id uuid.UUID) (*TakeHistory, error) {
	//TODO implement me
	panic("implement me")
}

func (t takeHistoryRepository) Add(newData *TakeHistory) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (t takeHistoryRepository) Update(newData *TakeHistory) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (t takeHistoryRepository) Delete(takeHistoryId uuid.UUID) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (t takeHistoryRepository) AddItem(item *TakeHistoryItem) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (t takeHistoryRepository) UpdateItem(item *TakeHistoryItem) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (t takeHistoryRepository) DeleteItem(takeHistoryItemId uuid.UUID) (bool, error) {
	//TODO implement me
	panic("implement me")
}
