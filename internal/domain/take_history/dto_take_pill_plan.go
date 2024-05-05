package takehistory

import (
	"github.com/google/uuid"
	"time"
)

type TakePlanRequest struct {
	UserId   uuid.UUID
	TakeDate time.Time
}
