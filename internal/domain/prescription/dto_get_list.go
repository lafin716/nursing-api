package prescription

import "github.com/google/uuid"

type GetListRequest struct {
	UserId uuid.UUID `json:"-"`
}

type GetListResponse struct {
	Success bool
	Message string
	Data    []*Prescription
	Error   error
}

func OkGetList(prescriptions []*Prescription) *GetListResponse {
	return &GetListResponse{
		Success: true,
		Message: "조회되었습니다.",
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
