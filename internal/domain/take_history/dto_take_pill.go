package takehistory

import "github.com/google/uuid"

type TakePillRequest struct {
	ItemId     uuid.UUID
	TakeAmount float64
}

type TakePillResponse struct {
	Success bool
	Message string
	Error   error
}

func OkTakePill() *TakePillResponse {
	return OkBaseTakePill("복용처리가 완료되었습니다.")
}

func OkUnTakePill() *TakePillResponse {
	return OkBaseTakePill("복용 미처리가 완료되었습니다.")
}

func OkBaseTakePill(message string) *TakePillResponse {
	return &TakePillResponse{
		Success: true,
		Message: message,
	}
}

func FailTakePill(message string, err error) *TakePillResponse {
	return &TakePillResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
