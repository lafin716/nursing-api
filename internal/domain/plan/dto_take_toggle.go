package plan

import "github.com/google/uuid"

type TakeToggleRequest struct {
	UserId     uuid.UUID
	TargetDate string    `json:"target_date" validate:"required"`
	TimezoneId uuid.UUID `json:"timezone_id" validate:"required"`
}
