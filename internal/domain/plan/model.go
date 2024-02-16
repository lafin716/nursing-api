package plan

import (
	"github.com/google/uuid"
)

type TakePlan struct {
	CurrentDate string `json:"current_date"`
	Plans       []Plan `json:"plans"`
}

type Plan struct {
	PlanName     string `json:"plan_name"`
	TakeTimeZone string `json:"take_time_zone"`
	TakeStatus   bool   `json:"take_status"`
	TakeDate     string `json:"take_date"`
	Pills        []Pill `json:"pills"`
	Memo         string `json:"memo"`
}

type Pill struct {
	PillName   string    `json:"pill_name"`
	MedicineId uuid.UUID `json:"medicine_id"`
	TakeStatus bool      `json:"take_status"`
	TakeDate   string    `json:"take_date"`
	TakeAmount float64   `json:"take_amount"`
}

type Summary struct {
	Year  string        `json:"year"`
	Month string        `json:"month"`
	Items []SummaryItem `json:"items"`
}

type SummaryItem struct {
	Date     string `json:"date"`
	IsExists bool   `json:"isExists"`
}

type UseCase interface {
	Add(req *AddPlanRequest) *AddPlanResponse
	GetByMonth(req *GetByMonthRequest) *GetByMonthResponse
	GetByDate(req *GetByDateRequest) *GetByDateResponse
	Take(req *TakeToggleRequest) *TakeToggleResponse
	PillToggle(req *PillToggleRequest) *PillToggleResponse
	UpdateMemo(req *UpdateMemoRequest) *UpdateMemoResponse
}
