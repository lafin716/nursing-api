package plan

import "github.com/google/uuid"

type UnTakePillRequest struct {
	UserId       uuid.UUID
	TakeTimeZone string    `json:"take_time_zone" validate:"required"`
	MedicineId   uuid.UUID `json:"medicine_id" validate:"required"`
}

type UnTakePillResponse struct {
	Success bool
	Message string
	Error   error
}

func OkUnTakePill() *UnTakePillResponse {
	return &UnTakePillResponse{
		Success: true,
		Message: "복용처리 되었습니다.",
	}
}

func FailUnTakePill(message string, err error) *UnTakePillResponse {
	return &UnTakePillResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
