package prescription

import "github.com/google/uuid"

type GetItemListRequest struct {
	UserId         uuid.UUID `json:"-"`
	PrescriptionId uuid.UUID `json:"prescription_id"`
}

type GetItemListResponse struct {
	Success bool
	Message string
	Data    []*PrescriptionItem
	Error   error
}

func OkGetItemListMessage(message string) *GetItemListResponse {
	return &GetItemListResponse{
		Success: true,
		Message: message,
		Data:    []*PrescriptionItem{},
	}
}

func OkGetItemList(prescriptions []*PrescriptionItem) *GetItemListResponse {
	return &GetItemListResponse{
		Success: true,
		Message: "처방전 의약품이 조회되었습니다.",
		Data:    prescriptions,
	}
}

func FailGetItemList(message string, err error) *GetItemListResponse {
	return &GetItemListResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
