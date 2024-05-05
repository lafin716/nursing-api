package takehistory

import (
	"github.com/google/uuid"
	"time"
)

type GetDetailRequest struct {
	UserId     uuid.UUID
	TargetDate time.Time
}
