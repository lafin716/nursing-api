package plan

import "github.com/google/uuid"

type PillToggleRequest struct {
	UserId uuid.UUID
	PillId uuid.UUID `json:"pill_id" validate:"required"`
}
