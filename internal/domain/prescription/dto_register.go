package prescription

import (
	"github.com/google/uuid"
)

type RegisterRequest struct {
	UserId           uuid.UUID             `json:"user_id" validate:"required"`
	PrescriptionName string                `json:"prescription_name" validate:"max=30"`
	HospitalName     string                `json:"hospital_name" validate:"required"`
	TakeDays         int                   `json:"take_days" validate:"required"`
	StartedAt        string                `json:"started_at" validate:"required"`
	FinishedAt       string                `json:"finished_at" validate:"required"`
	Memo             string                `json:"memo"`
	Items            []RegisterItemRequest `json:"items" validate:"required"`
}

type RegisterItemRequest struct {
	MedicineId   uuid.UUID `json:"medicine_id" validate:"required"`
	MedicineName string    `json:"medicine_name" validate:"required"`
	TakeTimeZone string    `json:"take_time_zone" validate:"required"`
	TakeMoment   string    `json:"take_moment"`
	TakeEtc      string    `json:"take_etc"`
	TakeAmount   float64   `json:"take_amount"`
	MedicineUnit string    `json:"medicine_unit"`
	Memo         string    `json:"memo"`
}

type RegisterResponse struct {
	Success bool
	Message string
	ID      uuid.UUID `json:"id"`
	Error   error
}

func OkRegister(id uuid.UUID) *RegisterResponse {
	return &RegisterResponse{
		Success: true,
		Message: "처방전이 등록되었습니다.",
		ID:      id,
	}
}

func FailRegister(message string, err error) *RegisterResponse {
	return &RegisterResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
