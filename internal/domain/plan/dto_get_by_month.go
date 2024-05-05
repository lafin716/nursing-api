package plan

import (
	"github.com/google/uuid"
)

type GetByMonthRequest struct {
	UserId uuid.UUID
	Year   string `json:"year" validate:"required"`
	Month  string `json:"month" validate:"required"`
}
