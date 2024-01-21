package plan

import (
	"github.com/google/uuid"
)

type GetByDateRequest struct {
	UserId      uuid.UUID
	CurrentDate string `query:"current_date"`
}

type GetByDateResponse struct {
	Success bool
	Message string
	Data    *TakePlan
	Error   error
}

func OkGetByDate(data *TakePlan) *GetByDateResponse {
	return &GetByDateResponse{
		Success: true,
		Message: "해당 날짜의 복용계획이 조회되었습니다.",
		Data:    data,
	}
}

func FailGetByDate(message string, err error) *GetByDateResponse {
	return &GetByDateResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
