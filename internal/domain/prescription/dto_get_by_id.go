package prescription

import (
	"github.com/google/uuid"
	"nursing_api/internal/domain/plan"
)

type GetByIdRequest struct {
	ID     uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"-"`
}

type GetByIdResponse struct {
	Success bool
	Message string
	Data    []*plan.Plan
	Error   error
}

func OkGetById(data []*plan.Plan) *GetByIdResponse {
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
