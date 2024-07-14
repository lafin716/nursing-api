package timezone

import "github.com/google/uuid"

type IsDeletableTimeZoneRequest struct {
	UserId     uuid.UUID `json:"-"`
	TimezoneId uuid.UUID `json:"timezone_id"`
}
