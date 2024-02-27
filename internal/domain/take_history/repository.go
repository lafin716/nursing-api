package takehistory

import (
	"context"
	"github.com/google/uuid"
	"log"
	"nursing_api/pkg/database"
	"nursing_api/pkg/ent"
	schema "nursing_api/pkg/ent/takehistory"
	schemaItem "nursing_api/pkg/ent/takehistoryitem"
	"time"
)

type Repository interface {
	GetList(userId uuid.UUID) ([]*TakeHistory, error)
	GetListByDate(userId uuid.UUID, today time.Time) ([]*TakeHistory, error)
	GetById(id uuid.UUID) (*TakeHistory, error)
	GetByTimezoneId(userId uuid.UUID, timezoneId uuid.UUID, date time.Time) (*TakeHistory, error)
	GetItemById(userId uuid.UUID, takeHistoryItemId uuid.UUID) (*TakeHistoryItem, error)
	GetItemsByHistoryId(userId uuid.UUID, historyId uuid.UUID, date time.Time) ([]*TakeHistoryItem, error)
	Add(newData *TakeHistory) (*TakeHistory, error)
	Update(newData *TakeHistory) (bool, error)
	Delete(takeHistoryId uuid.UUID) (bool, error)
	AddItem(item *TakeHistoryItem) (bool, error)
	UpdateItem(item *TakeHistoryItem) (bool, error)
	DeleteItem(takeHistoryItemId uuid.UUID) (bool, error)
	DeleteItemByHistoryId(takeHistoryId uuid.UUID) (bool, error)
}

type repository struct {
	root       *ent.Client
	client     *ent.TakeHistoryClient
	itemClient *ent.TakeHistoryItemClient
	ctx        context.Context
}

func NewRepository(
	dbClient *database.DatabaseClient,
) Repository {
	return &repository{
		root:       dbClient.Client,
		client:     dbClient.Client.TakeHistory,
		itemClient: dbClient.Client.TakeHistoryItem,
		ctx:        dbClient.Ctx,
	}
}
func (t *repository) GetList(userId uuid.UUID) ([]*TakeHistory, error) {
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

func (t *repository) GetByTimezoneId(userId uuid.UUID, timezoneId uuid.UUID, date time.Time) (*TakeHistory, error) {
	log.Println("takehistory.GetByTimezoneId")
	history, err := t.root.Debug().TakeHistory.
		Query().
		Where(
			schema.And(
				schema.UserID(userId),
				schema.TimezoneID(timezoneId),
				schema.TakeDateGTE(date),
				schema.TakeDateLT(date.AddDate(0, 0, 1)),
			),
		).Only(t.ctx)
	if err != nil {
		return nil, err
	}

	return toModel(history), nil
}

func (t *repository) GetItemById(userId uuid.UUID, takeHistoryItemId uuid.UUID) (*TakeHistoryItem, error) {
	item, err := t.itemClient.
		Query().
		Where(
			schemaItem.UserID(userId),
			schemaItem.ID(takeHistoryItemId),
		).Only(t.ctx)
	if err != nil {
		return nil, err
	}

	return toModelItem(item), nil
}

func (t *repository) GetItemsByHistoryId(userId uuid.UUID, historyId uuid.UUID, date time.Time) ([]*TakeHistoryItem, error) {
	historyItems, err := t.itemClient.Query().Where(
		schemaItem.And(
			schemaItem.UserID(userId),
			schemaItem.TakeHistoryID(historyId),
			schemaItem.TakeDateGTE(date),
			schemaItem.TakeDateLT(date.AddDate(0, 0, 1)),
		),
	).All(t.ctx)
	if err != nil {
		return nil, err
	}

	return toModelItemList(historyItems), nil
}

func (t *repository) GetListByDate(userId uuid.UUID, date time.Time) ([]*TakeHistory, error) {
	found, err := t.client.
		Query().
		Where(
			schema.UserID(userId),
			schema.TakeDateGTE(date),
			schema.TakeDateLT(date.AddDate(0, 0, 1)),
		).
		All(t.ctx)
	if err != nil {
		return nil, err
	}

	return toModelList(found), nil
}

func (t *repository) GetById(id uuid.UUID) (*TakeHistory, error) {
	found, err := t.client.
		Query().
		Where(schema.ID(id)).
		Only(t.ctx)
	if err != nil {
		return nil, err
	}

	return toModel(found), nil
}

func (t *repository) Add(newData *TakeHistory) (*TakeHistory, error) {
	saved, err := t.client.
		Create().
		SetUserID(newData.UserId).
		SetTimezoneID(newData.TimezoneId).
		SetTakeDate(newData.TakeDate).
		SetTakeStatus(string(newData.TakeStatus)).
		SetMemo(newData.Memo).
		SetCreatedAt(newData.CreatedAt).
		Save(t.ctx)
	if err != nil {
		return nil, err
	}

	return toModel(saved), nil
}

func (t *repository) Update(newData *TakeHistory) (bool, error) {
	result, err := t.client.
		Update().
		SetTakeDate(newData.TakeDate).
		SetTakeStatus(string(newData.TakeStatus)).
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

func (t *repository) Delete(takeHistoryId uuid.UUID) (bool, error) {
	result, err := t.client.Delete().Where(schema.ID(takeHistoryId)).Exec(t.ctx)
	if err != nil {
		return false, err
	}

	return result > 0, nil
}

func (t *repository) AddItem(item *TakeHistoryItem) (bool, error) {
	_, err := t.itemClient.
		Create().
		SetUserID(item.UserId).
		SetTakeHistoryID(item.TakeHistoryId).
		SetPrescriptionItemID(item.PrescriptionItemId).
		SetTakeAmount(item.TakeAmount).
		SetTakeUnit(item.TakeUnit).
		SetTakeDate(item.TakeDate).
		SetTakeStatus(string(item.TakeStatus)).
		SetCreatedAt(time.Now()).
		Save(t.ctx)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (t *repository) UpdateItem(item *TakeHistoryItem) (bool, error) {
	err := t.itemClient.
		Update().
		SetTakeStatus(string(item.TakeStatus)).
		SetMemo(item.Memo).
		Where(
			schemaItem.ID(item.ID),
			schemaItem.UserID(item.UserId),
		).
		Exec(t.ctx)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (t *repository) DeleteItem(takeHistoryItemId uuid.UUID) (bool, error) {
	result, err := t.itemClient.Delete().Where(schemaItem.ID(takeHistoryItemId)).Exec(t.ctx)
	if err != nil {
		return false, err
	}

	return result > 0, nil
}

func (t *repository) DeleteItemByHistoryId(takeHistoryId uuid.UUID) (bool, error) {
	result, err := t.itemClient.Delete().Where(schemaItem.TakeHistoryID(takeHistoryId)).Exec(t.ctx)
	if err != nil {
		return false, err
	}

	return result > 0, nil
}
