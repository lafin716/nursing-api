package timezone

import (
	"github.com/google/uuid"
)

type UpdateTimeZoneRequest struct {
	ID       uuid.UUID `json:"id" validate:"required"`
	UserId   uuid.UUID `json:"-"`
	Name     string    `json:"name"`
	UseAlert *bool     `json:"use_alert"`
	Midday   string    `json:"midday"`
	Hour     string    `json:"hour"`
	Minute   string    `json:"minute"`
}

type UpdateTimeZoneResponse struct {
	Success bool
	Message string
	Data    *TimeZone
	Error   error
}

func OkUpdateTimeZone(data *TimeZone) *UpdateTimeZoneResponse {
	return &UpdateTimeZoneResponse{
		Success: true,
		Message: "정상적으로 업데이트 되었습니다",
		Data:    data,
	}
}

func FailUpdateTimeZone(message string, err error) *UpdateTimeZoneResponse {
	return &UpdateTimeZoneResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
