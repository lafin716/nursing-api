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
	PlanName   string    `json:"plan_name"`
	Timezone   string    `json:"time_zone"`
	TakeDate   string    `json:"take_date"`
	Pills      []Pill    `json:"pills"`

	// 복용히스토리
	TakeStatus bool   `json:"take_status"`
	Memo       string `json:"memo"`
}

type Pill struct {
	// 복용계획
	PillName   string    `json:"pill_name"`
	MedicineId uuid.UUID `json:"medicine_id"`
	TakeUnit   string    `json:"take_unit"`
	TakeAmount float64   `json:"take_amount"`

	// 복용히스토리
	TakeStatus bool   `json:"take_status"`
	TakeDate   string `json:"take_date"`
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
