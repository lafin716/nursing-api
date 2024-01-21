package plan

import (
	"github.com/google/uuid"
)

type TakePlan struct {
	CurrentDate string     `json:"current_date"`
	Plans       []PlanItem `json:"plans"`
}

type PlanItem struct {
	PlanName     string     `json:"plan_name"`
	TakeTimeZone string     `json:"take_time_zone"`
	TakeStatus   bool       `json:"take_status"`
	TakeDate     string     `json:"take_date"`
	Pills        []PillItem `json:"pills"`
	Memo         string     `json:"memo"`
}

type PillItem struct {
	PillName   string    `json:"pill_name"`
	MedicineId uuid.UUID `json:"medicine_id"`
	TakeStatus bool      `json:"take_status"`
	TakeDate   string    `json:"take_date"`
	TakeAmount float64   `json:"take_amount"`
}

type PlanSummary struct {
	Year  string            `json:"year"`
	Month string            `json:"month"`
	Items []PlanSummaryItem `json:"items"`
}

type PlanSummaryItem struct {
	Date     string `json:"date"`
	IsExists bool   `json:"isExists"`
}

type PlanUseCase interface {
	GetByMonth(req *GetByMonthRequest) *GetByMonthResponse
	GetByDate(req *GetByDateRequest) *GetByDateResponse
	TakePlan(req *TakePlanRequest) *TakePlanResponse
	UnTakePlan(req *UnTakePlanRequest) *UnTakePlanResponse
	TakePill(req *TakePillRequest) *TakePillResponse
	UnTakePill(req *UnTakePillRequest) *UnTakePillResponse
	UpdateMemo(req *UpdateMemoRequest) *UpdateMemoResponse
}
