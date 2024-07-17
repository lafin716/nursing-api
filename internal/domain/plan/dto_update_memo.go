package plan

import "github.com/google/uuid"

type UpdateMemoRequest struct {
	UserId     uuid.UUID
	TimezoneId uuid.UUID `json:"timezone_id" validate:"required"`
	Date       string    `json:"date" validate:"required"`
	Memo       string    `json:"memo"`
}
