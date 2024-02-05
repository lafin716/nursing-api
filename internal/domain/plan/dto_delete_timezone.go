package plan

import "github.com/google/uuid"

type DeleteTimeZoneRequest struct {
	ID     uuid.UUID `json:"id" validate:"required"`
	UserId uuid.UUID `json:"-"`
}

type DeleteTimeZoneResponse struct {
	Success bool
	Message string
	Error   error
}

func OkDeleteTimeZone() *DeleteTimeZoneResponse {
	return &DeleteTimeZoneResponse{
		Success: true,
		Message: "시간대가 삭제되었습니다",
	}
}

func FailDeleteTimeZone(message string, err error) *DeleteTimeZoneResponse {
	return &DeleteTimeZoneResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
