package timezone

import "github.com/google/uuid"

type DeleteTimeZoneRequest struct {
	ID     uuid.UUID `json:"id" validate:"required"`
	UserId uuid.UUID `json:"-"`
}
