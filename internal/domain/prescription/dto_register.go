package prescription

import (
	"github.com/google/uuid"
	"time"
)

type RegisterPrescription struct {
	UserId           uuid.UUID `json:"user_id" validate:"required"`
	PrescriptionName string    `json:"prescription_name" validate:"required,max=30"`
	HospitalName     string    `json:"hospital_name" validate:"required"`
	TakeDays         int       `json:"take_days" validate:"required"`
	StartedAt        time.Time `json:"started_at" validate:"required"`
	FinishedAt       time.Time `json:"finished_at" validate:"required"`
	Memo             string    `json:"memo" validate:"required"`
}
