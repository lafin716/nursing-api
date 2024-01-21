package plan

import (
	"github.com/google/uuid"
)

type GetByMonthRequest struct {
	UserId uuid.UUID
	Year   string `json:"year" validate:"required"`
	Month  string `json:"month" validate:"required"`
}

type GetByMonthResponse struct {
	Success bool
	Message string
	Data    *PlanSummary
	Error   error
}

func OkGetByMonth(data *PlanSummary) *GetByMonthResponse {
	return &GetByMonthResponse{
		Success: true,
		Message: "해당 월의 복용계획이 조회되었습니다.",
		Data:    data,
	}
}

func FailGetByMonth(message string, err error) *GetByMonthResponse {
	return &GetByMonthResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
