package prescription

import "github.com/google/uuid"

type DeleteItemRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

type DeleteItemResponse struct {
	Success bool
	Message string
	Error   error
}

func OkDeleteItem() *DeleteItemResponse {
	return &DeleteItemResponse{
		Success: true,
		Message: "처방전 의약품이 삭제되었습니다.",
	}
}

func FailDeleteItem(message string, err error) *DeleteItemResponse {
	return &DeleteItemResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
