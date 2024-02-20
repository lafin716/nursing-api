package prescription

import (
	"github.com/google/uuid"
)

type GetByIdRequest struct {
	UserId         uuid.UUID `json:"-"`
	PrescriptionId uuid.UUID `json:"prescription_id"`
}

type GetByIdResponse struct {
	Success bool
	Message string
	Data    *Prescription
	Error   error
}

func OkGetById(data *Prescription) *GetByIdResponse {
	return &GetByIdResponse{
		Success: true,
		Message: "처방전이 조회 되었습니다.",
		Data:    data,
	}
}

func FailGetById(message string, err error) *GetByIdResponse {
	return &GetByIdResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
