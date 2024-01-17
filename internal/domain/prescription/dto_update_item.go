package prescription

import "github.com/google/uuid"

type UpdateItemRequest struct {
	ID           uuid.UUID `json:"id" validate:"required"`
	MedicineId   uuid.UUID `json:"medicine_id"`
	MedicineName string    `json:"medicine_name"`
	TakeTimeZone string    `json:"take_time_zone"`
	TakeMoment   string    `json:"take_moment"`
	TakeEtc      string    `json:"take_etc"`
	TakeAmount   float64   `json:"take_amount"`
	MedicineUnit string    `json:"medicine_unit"`
	Memo         string    `json:"memo"`
}

type UpdateItemResponse struct {
	Success bool
	Message string
	Error   error
}

func OkUpdateItem() *UpdateItemResponse {
	return &UpdateItemResponse{
		Success: true,
		Message: "처방전 의약품이 업데이트되었습니다.",
	}
}

func FailUpdateItem(message string, err error) *UpdateItemResponse {
	return &UpdateItemResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
