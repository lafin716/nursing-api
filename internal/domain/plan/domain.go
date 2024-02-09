package plan

import (
	"github.com/google/uuid"
	"time"
)

type TimeZone struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Name      string    `json:"name"`
	IsDefault bool      `json:"is_default"`
	UseAlert  bool      `json:"use_alert"`
	Meridiem  string    `json:"meridiem"`
	Hour      string    `json:"hour"`
	Minute    string    `json:"minute"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TimeZoneLink struct {
	ID             uuid.UUID `json:"id"`
	PrescriptionId uuid.UUID `json:"prescription_id"`
	TimeZoneId     uuid.UUID `json:"timezone_id"`
	TimeZoneName   string    `json:"timezone_name"`
	UseAlert       bool      `json:"use_alert"`
	Meridiem       string    `json:"meridiem"`
	Hour           string    `json:"hour"`
	Minute         string    `json:"minute"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

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

type PlanRepository interface {
	// 계획 시간대 템플릿
	GetTimeZones(userId uuid.UUID) ([]*TimeZone, error)
	GetTimeZone(id uuid.UUID, userId uuid.UUID) (*TimeZone, error)
	CreateTimeZone(model *TimeZone) (*TimeZone, error)
	UpdateTimeZone(model *TimeZone) (bool, error)
	DeleteTimeZone(id uuid.UUID, userId uuid.UUID) (bool, error)
	GetDuplicate(userId uuid.UUID, name string, meridiem string, hour string, minute string) (*TimeZone, error)
}

// 복용계획 시간대 로직
type TimeZoneUseCase interface {
	GetList(userId uuid.UUID) ([]*TimeZone, error)
	Create(req *CreateTimeZoneRequest) *CreateTimeZoneResponse
	Update(req *UpdateTimeZoneRequest) *UpdateTimeZoneResponse
	Delete(req *DeleteTimeZoneRequest) *DeleteTimeZoneResponse
}

// 복용계획 로직
type PlanUseCase interface {
	GetByMonth(req *GetByMonthRequest) *GetByMonthResponse
	GetByDate(req *GetByDateRequest) *GetByDateResponse
	TakePlan(req *TakePlanRequest) *TakePlanResponse
	UnTakePlan(req *UnTakePlanRequest) *UnTakePlanResponse
	TakePill(req *TakePillRequest) *TakePillResponse
	UnTakePill(req *UnTakePillRequest) *UnTakePillResponse
	UpdateMemo(req *UpdateMemoRequest) *UpdateMemoResponse
}
