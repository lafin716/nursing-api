package prescription

import "github.com/google/uuid"

type DeleteRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

type DeleteResponse struct {
	Success bool
	Message string
	Error   error
}

func OkDelete() *DeleteResponse {
	return &DeleteResponse{
		Success: true,
		Message: "처방전이 삭제되었습니다.",
	}
}

func FailDelete(message string, err error) *DeleteResponse {
	return &DeleteResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
