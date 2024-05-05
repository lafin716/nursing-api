package plan

import (
	"github.com/google/uuid"
)

type GetByDateRequest struct {
	UserId      uuid.UUID
	CurrentDate string `query:"date"`
}
