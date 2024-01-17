package prescription

import "github.com/google/uuid"

type AddItemRequest struct {
	PrescriptionId uuid.UUID `json:"prescription_id" validate:"required"`
	MedicineId     uuid.UUID `json:"medicine_id" validate:"required"`
	MedicineName   string    `json:"medicine_name" validate:"required,min=1"`
	TakeTimeZone   string    `json:"take_time_zone" validate:"required,min=1"`
	TakeMoment     string    `json:"take_moment"`
	TakeEtc        string    `json:"take_etc"`
	TakeAmount     float64   `json:"take_amount"`
	MedicineUnit   string    `json:"medicine_unit"`
	Memo           string    `json:"memo"`
}

type AddItemResponse struct {
	Success bool
	Message string
	Error   error
}

func OkAddItem() *AddItemResponse {
	return &AddItemResponse{
		Success: true,
		Message: "처방전 의약품이 추가되었습니다.",
	}
}

func FailAddItem(message string, err error) *AddItemResponse {
	return &AddItemResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
