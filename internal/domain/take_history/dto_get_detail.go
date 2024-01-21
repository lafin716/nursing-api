package takehistory

import (
	"github.com/google/uuid"
	"time"
)

type GetDetailRequest struct {
	UserId     uuid.UUID
	TargetDate time.Time
}

type GetDetailResponse struct {
	Success bool
	Message string
	Data    *TakeHistory
	Error   error
}

func OkGetDetail(data *TakeHistory) *GetDetailResponse {
	return &GetDetailResponse{
		Success: true,
		Message: "복용내역 상세데이터가 조회되었습니다",
		Data:    data,
	}
}

func FailGetDetail(message string, err error) *GetDetailResponse {
	return &GetDetailResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
