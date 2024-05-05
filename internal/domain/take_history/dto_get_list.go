package takehistory

import (
	"github.com/google/uuid"
)

type GetListRequest struct {
	UserId     uuid.UUID
	TargetDate string `query:"date"`
}
