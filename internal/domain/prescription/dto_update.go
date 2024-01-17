package prescription

import "github.com/google/uuid"

type UpdateRequest struct {
	ID               uuid.UUID `json:"id" validate:"required"`
	HospitalName     string    `json:"hospital_name" validate:"max=30"`
	PrescriptionName string    `json:"prescription_name"`
	TakeDays         int       `json:"take_days"`
	StartedAt        string    `json:"started_at"`
	FinishedAt       string    `json:"finished_at"`
	Memo             string    `json:"memo"`
}

type UpdateResponse struct {
	Success bool
	Message string
	Data    *Prescription
	Error   error
}

func OkUpdate(data *Prescription) *UpdateResponse {
	return &UpdateResponse{
		Success: true,
		Message: "처방전이 업데이트되었습니다.",
		Data:    data,
	}
}

func FailUpdate(message string, err error) *UpdateResponse {
	return &UpdateResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
