package plan

import "github.com/google/uuid"

type AddPlanRequest struct {
	UserId    uuid.UUID            `json:"-"`
	Name      string               `json:"name" validate:"required"`
	Hospital  string               `json:"hospital" validate:"required"`
	StartedAt string               `json:"started_at" validate:"required"`
	TakeDays  int                  `json:"take_days" validate:"required"`
	Memo      string               `json:"memo"`
	Timezones []AddTimezoneRequest `json:"timezones"`
}

type AddTimezoneRequest struct {
	TimezoneId uuid.UUID            `json:"timezone_id" validate:"required"`
	Medicines  []AddMedicineRequest `json:"medicines"`
}

type AddMedicineRequest struct {
	MedicineId uuid.UUID `json:"medicine_id"`
	Name       string    `json:"name"`
	TakeAmount float64   `json:"take_amount"`
	TakeUnit   string    `json:"take_unit"`
	Memo       string    `json:"memo"`
}

type AddPlanResponse struct {
	Success bool
	Message string
	Error   error
}

func OkAddPlan() *AddPlanResponse {
	return &AddPlanResponse{
		Success: true,
		Message: "복용계획이 정상 추가되었습니다",
	}
}

func FailAddPlan(message string, err error) *AddPlanResponse {
	return &AddPlanResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
