package plan

import "github.com/google/uuid"

type TakeToggleRequest struct {
	UserId         uuid.UUID
	PlanTimezoneId uuid.UUID `json:"plan_timezone_id" validate:"required"`
}

type TakeToggleResponse struct {
	Success bool
	Message string
	Error   error
}

func OkTakeToggle(toggle bool) *TakeToggleResponse {
	msg := "복용처리 되었습니다"
	if !toggle {
		msg = "미복용처리 되었습니다"
	}
	return &TakeToggleResponse{
		Success: true,
		Message: msg,
	}
}

func FailTakeToggle(message string, err error) *TakeToggleResponse {
	return &TakeToggleResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
