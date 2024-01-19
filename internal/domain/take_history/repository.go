package take_history

import (
	"context"
	"github.com/google/uuid"
	"nursing_api/pkg/database"
	"nursing_api/pkg/ent"
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
func (t takeHistoryRepository) GetList() ([]*TakeHistory, error) {
	//TODO implement me
	panic("implement me")
}

func (t takeHistoryRepository) GetByToday(userId uuid.UUID, today time.Time) (*TakeHistory, error) {
	//TODO implement me
	panic("implement me")
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
