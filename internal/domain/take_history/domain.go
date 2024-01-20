package takehistory

import (
	"github.com/google/uuid"
	"time"
)

type TakeHistory struct {
	ID               uuid.UUID          `json:"id"`
	UserId           uuid.UUID          `json:"user_id"`
	PrescriptionId   uuid.UUID          `json:"prescription_id"`
	TakeDate         time.Time          `json:"take_date"`
	TakeStatus       string             `json:"take_status"`
	Memo             string             `json:"memo"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
	TakeHistoryItems []*TakeHistoryItem `json:"items"`
}

type TakeHistoryItem struct {
	ID                 uuid.UUID `json:"id"`
	UserId             uuid.UUID `json:"user_id"`
	TakeHistoryId      uuid.UUID `json:"take_history_id"`
	PrescriptionItemId uuid.UUID `json:"prescription_item_id"`
	TakeStatus         string    `json:"take_status"`
	TakeAmount         float64   `json:"take_amount"`
	TakeTimeZone       string    `json:"take_time_zone"`
	TakeMoment         string    `json:"take_moment"`
	TakeEtc            string    `json:"take_etc"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type TakeHistoryRepository interface {
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

type TakeHistoryUseCase interface {
	GetList(req *GetListRequest) *GetListResponse
	GetDetail(req *GetDetailRequest) *GetDetailResponse
	TakePlanToggle(req *TakePlanRequest) *TakePlanResponse
	TakePillToggle(req *TakePillRequest) *TakePillResponse
}
