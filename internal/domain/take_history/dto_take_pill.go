package takehistory

import "github.com/google/uuid"

type TakePillRequest struct {
	ItemId     uuid.UUID
	TakeAmount float64
}
