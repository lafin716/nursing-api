package timezone

import "github.com/google/uuid"

type IsDeletableTimeZoneRequest struct {
	ID     uuid.UUID `json:"id" validate:"required"`
	UserId uuid.UUID `json:"-"`
}
