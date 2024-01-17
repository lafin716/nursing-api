package prescription

import (
	"github.com/google/uuid"
	"time"
)

type GetByDateRequest struct {
	UserId uuid.UUID
	Date   time.Time
}

type GetByDateResponse struct {
	Success bool
	Message string
	Data    interface{}
	Error   error
}

func OkGetByDate(data interface{}) *GetByDateResponse {
	return &GetByDateResponse{
		Success: true,
		Message: "복용 플랜이 조회 되었습니다.",
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
