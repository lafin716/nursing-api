package plan

import "github.com/google/uuid"

type PillToggleRequest struct {
	UserId uuid.UUID
	PillId uuid.UUID `json:"pill_id" validate:"required"`
}

type PillToggleResponse struct {
	Success bool
	Message string
	Error   error
}

func OkPillToggle(toggle bool) *PillToggleResponse {
	msg := "복용처리 되었습니다"
	if !toggle {
		msg = "미복용처리 되었습니다"
	}
	return &PillToggleResponse{
		Success: true,
		Message: msg,
	}
}

func FailPillToggle(message string, err error) *PillToggleResponse {
	return &PillToggleResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
