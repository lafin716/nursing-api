package plan

import "github.com/google/uuid"

type TakePillRequest struct {
	UserId       uuid.UUID
	TakeTimeZone string    `json:"take_time_zone" validate:"required"`
	MedicineId   uuid.UUID `json:"medicine_id" validate:"required"`
}

type TakePillResponse struct {
	Success bool
	Message string
	Error   error
}

func OkTakePill() *TakePillResponse {
	return &TakePillResponse{
		Success: true,
		Message: "복용처리 되었습니다.",
	}
}

func FailTakePill(message string, err error) *TakePillResponse {
	return &TakePillResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
