package timezone

import (
	"github.com/google/uuid"
)

type CreateTimeZoneRequest struct {
	UserId   uuid.UUID `json:"-"`
	Name     string    `json:"name" validate:"required"`
	UseAlert *bool     `json:"use_alert" validate:"required"`
	Midday   string    `json:"midday" validate:"required,max=2"`
	Hour     string    `json:"hour" validate:"required,max=2"`
	Minute   string    `json:"minute" validate:"required,max=2"`
}
