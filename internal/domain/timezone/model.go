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

type Repository interface {
	GetTimeZones(userId uuid.UUID) ([]*TimeZone, error)
	GetTimeZone(id uuid.UUID, userId uuid.UUID) (*TimeZone, error)
	CreateTimeZone(model *TimeZone) (*TimeZone, error)
	UpdateTimeZone(model *TimeZone) (bool, error)
	DeleteTimeZone(id uuid.UUID, userId uuid.UUID) (bool, error)
	GetDuplicate(userId uuid.UUID, name string, midday string, hour string, minute string) (*TimeZone, error)
}

// 복용계획 시간대 로직
type UseCase interface {
	GetList(userId uuid.UUID) ([]*TimeZone, error)
	Create(req *CreateTimeZoneRequest) *CreateTimeZoneResponse
	Update(req *UpdateTimeZoneRequest) *UpdateTimeZoneResponse
	Delete(req *DeleteTimeZoneRequest) *DeleteTimeZoneResponse
}
