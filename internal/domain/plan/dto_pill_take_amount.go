package plan

import "github.com/google/uuid"

type PillTakeAmountUpdateRequest struct {
	PrescriptionItemId uuid.UUID `json:"prescription_item_id" validate:"required"`
	TakeAmount         float64   `json:"take_amount" validate:"required"`
}
