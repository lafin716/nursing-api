package takehistory

import (
	"github.com/google/uuid"
	"time"
)

type TakeHistory struct {
	ID               uuid.UUID          `json:"id"`
	UserId           uuid.UUID          `json:"user_id"`
	PrescriptionId   uuid.UUID          `json:"prescription_id"`
	TimezoneId       uuid.UUID          `json:"timezone_id"`
	TakeDate         time.Time          `json:"take_date"`
	TakeStatus       TakeStatus         `json:"take_status"`
	Memo             string             `json:"memo"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
	TakeHistoryItems []*TakeHistoryItem `json:"items"`
}

type TakeHistoryItem struct {
	ID                 uuid.UUID      `json:"id"`
	UserId             uuid.UUID      `json:"user_id"`
	TakeHistoryId      uuid.UUID      `json:"take_history_id"`
	PrescriptionItemId uuid.UUID      `json:"prescription_item_id"`
	TakeStatus         TakePillStatus `json:"take_status"`
	TakeAmount         float64        `json:"take_amount"`
	TakeUnit           string         `json:"take_unit"`
	TakeDate           time.Time      `json:"take_date"`
	Memo               string         `json:"memo"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
}
