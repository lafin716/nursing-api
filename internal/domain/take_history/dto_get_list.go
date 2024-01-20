package takehistory

import (
	"github.com/google/uuid"
	"time"
)

type GetListRequest struct {
	userId     uuid.UUID
	targetDate time.Time
}

type GetListResponse struct {
	Success bool
	Message string
	Data    []*TakeHistory
	Error   error
}

func OkGetList(data []*TakeHistory) *GetListResponse {
	return &GetListResponse{
		Success: true,
		Message: "복용내역이 조회되었습니다",
		Data:    data,
	}
}

func FailGetList(message string, err error) *GetListResponse {
	return &GetListResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
