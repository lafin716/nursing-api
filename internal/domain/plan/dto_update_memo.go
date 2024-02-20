package plan

import "github.com/google/uuid"

type UpdateMemoRequest struct {
	UserId     uuid.UUID
	TimezoneId uuid.UUID `json:"timezone_id" validate:"required"`
	Date       string    `json:"date" validate:"required"`
	Memo       string    `json:"memo" validate:"required"`
}

type UpdateMemoResponse struct {
	Success bool
	Message string
	Error   error
}

func OkUpdateMemo() *UpdateMemoResponse {
	return &UpdateMemoResponse{
		Success: true,
		Message: "메모가 업데이트 되었습니다.",
	}
}

func FailUpdateMemo(message string, err error) *UpdateMemoResponse {
	return &UpdateMemoResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
