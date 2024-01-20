package takehistory

import (
	"github.com/google/uuid"
	"time"
)

type TakePlanRequest struct {
	UserId   uuid.UUID
	TakeDate time.Time
}

type TakePlanResponse struct {
	Success bool
	Message string
	Error   error
}

func OkTake() *TakePlanResponse {
	return OkTakePillPlan("복용처리가 완료되었습니다.")
}

func OkUnTake() *TakePlanResponse {
	return OkTakePillPlan("복용 미처리가 완료되었습니다.")
}

func OkTakePillPlan(message string) *TakePlanResponse {
	return &TakePlanResponse{
		Success: true,
		Message: message,
	}
}

func FailTakePillPlan(message string, err error) *TakePlanResponse {
	return &TakePlanResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
