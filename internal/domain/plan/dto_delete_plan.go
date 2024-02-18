package plan

import "github.com/google/uuid"

type DeletePlanRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

type DeletePlanResponse struct {
	Success bool
	Message string
	Error   error
}

func OkDeletePlan() *DeletePlanResponse {
	return &DeletePlanResponse{
		Success: true,
		Message: "복용계획이 삭제되었습니다",
	}
}

func FailDeletePlan(message string, err error) *DeletePlanResponse {
	return &DeletePlanResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
