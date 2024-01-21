package plan

import "github.com/google/uuid"

type UpdateMemoRequest struct {
	UserId       uuid.UUID
	TakeTimeZone string `json:"take_time_zone" validate:"required"`
	Memo         string `json:"memo" validate:"required"`
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
