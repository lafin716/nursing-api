package plan

import "github.com/google/uuid"

type DeletePlanRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}
