package timezone

import (
	"github.com/google/uuid"
)

type CreateTimeZoneRequest struct {
	UserId   uuid.UUID `json:"-"`
	Name     string    `json:"name" validate:"required"`
	UseAlert *bool     `json:"use_alert" validate:"required"`
	Midday   string    `json:"midday" validate:"required,max=2"`
	Hour     string    `json:"hour" validate:"required,max=2"`
	Minute   string    `json:"minute" validate:"required,max=2"`
}

type CreateTimeZoneResponse struct {
	Success bool
	Message string
	Data    *TimeZone
	Error   error
}

func OkCreateTimeZone(data *TimeZone) *CreateTimeZoneResponse {
	return &CreateTimeZoneResponse{
		Success: true,
		Message: "정상적으로 생성되었습니다",
		Data:    data,
	}
}

func FailCreateTimeZone(message string, err error) *CreateTimeZoneResponse {
	return &CreateTimeZoneResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
