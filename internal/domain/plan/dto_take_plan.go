package plan

import "github.com/google/uuid"

type TakePlanRequest struct {
	UserId       uuid.UUID
	TakeTimeZone string `json:"take_time_zone" validate:"required"`
}

type TakePlanResponse struct {
	Success bool
	Message string
	Error   error
}

func OkTakePlan() *TakePlanResponse {
	return &TakePlanResponse{
		Success: true,
		Message: "복용처리 되었습니다.",
	}
}

func FailTakePlan(message string, err error) *TakePlanResponse {
	return &TakePlanResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
