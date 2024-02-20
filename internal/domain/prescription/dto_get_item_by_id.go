package prescription

import "github.com/google/uuid"

type GetItemByIdRequest struct {
	UserId             uuid.UUID `json:"-"`
	PrescriptionItemId uuid.UUID `json:"id"`
}

type GetItemByIdResponse struct {
	Success bool
	Message string
	Data    *PrescriptionItem
	Error   error
}

func OkGetItemByIdMessage(message string) *GetItemByIdResponse {
	return &GetItemByIdResponse{
		Success: true,
		Message: message,
	}
}

func OkGetItemById(prescriptions *PrescriptionItem) *GetItemByIdResponse {
	return &GetItemByIdResponse{
		Success: true,
		Message: "처방전 의약품이 조회되었습니다.",
		Data:    prescriptions,
	}
}

func FailGetItemById(message string, err error) *GetItemByIdResponse {
	return &GetItemByIdResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
