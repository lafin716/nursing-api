package timezone

import (
	"github.com/google/uuid"
	"time"
)

type TimeZone struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Name      string    `json:"name"`
	IsDefault bool      `json:"is_default"`
	Midday    string    `json:"midday"`
	Hour      string    `json:"hour"`
	Minute    string    `json:"minute"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
