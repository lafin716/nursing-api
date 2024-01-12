package prescription

import (
	"github.com/google/uuid"
	"time"
)

type RegisterRequest struct {
	UserId           uuid.UUID             `json:"user_id" validate:"required"`
	PrescriptionName string                `json:"prescription_name" validate:"required,max=30"`
	HospitalName     string                `json:"hospital_name" validate:"required"`
	TakeDays         int                   `json:"take_days" validate:"required"`
	StartedAt        time.Time             `json:"started_at" validate:"required"`
	FinishedAt       time.Time             `json:"finished_at" validate:"required"`
	Memo             string                `json:"memo" validate:"required"`
	Items            []RegisterItemRequest `json:"items" validate:"required"`
}

type RegisterItemRequest struct {
	ID             uuid.UUID `json:"id" validate:"required"`
	PrescriptionId uuid.UUID `json:"prescription_id"`
	UserId         uuid.UUID `json:"user_id"`
	MedicineName   string    `json:"medicine_name" validate:"required"`
	TakeTimeZone   string    `json:"take_time_zone" validate:"required"`
	TakeMoment     string    `json:"take_moment"`
	TakeEtc        string    `json:"take_etc"`
	TakeAmount     float32   `json:"take_amount"`
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
		Message: "의약품 검색 완료",
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
