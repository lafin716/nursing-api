package timezone

import (
	"github.com/google/uuid"
)

type UpdateTimeZoneRequest struct {
	ID       uuid.UUID `json:"id" validate:"required"`
	UserId   uuid.UUID `json:"-"`
	Name     string    `json:"name"`
	UseAlert *bool     `json:"use_alert"`
	Midday   string    `json:"midday"`
	Hour     string    `json:"hour"`
	Minute   string    `json:"minute"`
}
