package prescription

import "github.com/google/uuid"

type GetListRequest struct {
	UserId uuid.UUID `json:"-"`
	Date   string    `json:"date"`
	Limit  int       `json:"limit"`
}

type GetListResponse struct {
	Success bool
	Message string
	Data    []*Prescription
	Error   error
}

func OkGetListMessage(message string) *GetListResponse {
	return &GetListResponse{
		Success: true,
		Message: message,
		Data:    []*Prescription{},
	}
}

func OkGetList(prescriptions []*Prescription) *GetListResponse {
	return &GetListResponse{
		Success: true,
		Message: "처방전 목록이 조회되었습니다.",
		Data:    prescriptions,
	}
}

func FailGetList(message string, err error) *GetListResponse {
	return &GetListResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
