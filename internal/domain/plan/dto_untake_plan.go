package plan

import "github.com/google/uuid"

type UnTakePlanRequest struct {
	UserId       uuid.UUID
	TakeTimeZone string `json:"take_time_zone" validate:"required"`
}

type UnTakePlanResponse struct {
	Success bool
	Message string
	Error   error
}

func OkUnTakePlan() *UnTakePlanResponse {
	return &UnTakePlanResponse{
		Success: true,
		Message: "미복용처리 되었습니다.",
	}
}

func FailUnTakePlan(message string, err error) *UnTakePlanResponse {
	return &UnTakePlanResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
