package timezonelink

import (
	"github.com/google/uuid"
	"time"
)

type TimeZoneLink struct {
	ID             uuid.UUID `json:"id"`
	PrescriptionId uuid.UUID `json:"prescription_id"`
	TimeZoneId     uuid.UUID `json:"timezone_id"`
	TimeZoneName   string    `json:"timezone_name"`
	UseAlert       bool      `json:"use_alert"`
	Midday         string    `json:"midday"`
	Hour           string    `json:"hour"`
	Minute         string    `json:"minute"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
