package plan

import (
	"github.com/google/uuid"
)

type TakePlan struct {
	CurrentDate string `json:"current_date"`
	Plans       []Plan `json:"plans"`
}

type Plan struct {
	// 복용계획
	TimezoneId uuid.UUID `json:"timezone_id"`
	Hour       int       `json:"hour"`
	Minute     int       `json:"minute"`
	PlanName   string    `json:"plan_name"`
	Timezone   string    `json:"time_zone"`
	Pills      []Pill    `json:"pills"`

	// 복용히스토리
	TakeStatus bool   `json:"take_status"`
	TakeDate   string `json:"take_date,omitempty"`
	Memo       string `json:"memo"`
}

type Pill struct {
	// 복용계획
	PrescriptionItemId uuid.UUID `json:"prescription_item_id"`
	PillName           string    `json:"pill_name"`
	MedicineId         uuid.UUID `json:"medicine_id"`
	TakeUnit           string    `json:"take_unit"`
	TakeAmount         float64   `json:"take_amount"`
	RemainAmount       float64   `json:"remain_amount"`
	TotalAmount        float64   `json:"total_amount"`

	// 복용히스토리
	TakeHistoryItemId uuid.UUID `json:"take_history_item_id,omitempty"`
	TakeStatus        bool      `json:"take_status"`
	TakeDate          string    `json:"take_date,omitempty"`
}

type Summary struct {
	Year  string        `json:"year"`
	Month string        `json:"month"`
	Items []SummaryItem `json:"items"`
}

type SummaryItem struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	StartedAt  string    `json:"started_at"`
	FinishedAt string    `json:"finished_at"`
	TakeDays   int       `json:"take_days"`
	IsExists   bool      `json:"isExists"`
}
