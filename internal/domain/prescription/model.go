package prescription

import (
	"github.com/google/uuid"
	"nursing_api/internal/domain/medicine"
	"time"
)

type Prescription struct {
	ID                uuid.UUID           `json:"id"`
	UserId            uuid.UUID           `json:"user_id"`
	PrescriptionName  string              `json:"name"`
	HospitalName      string              `json:"hospital"`
	TakeDays          int                 `json:"take_days"`
	StartedAt         time.Time           `json:"started_at"`
	FinishedAt        time.Time           `json:"finished_at"`
	CreatedAt         time.Time           `json:"created_at"`
	UpdatedAt         time.Time           `json:"updated_at"`
	PrescriptionItems []*PrescriptionItem `json:"prescription_items"`
}

type PrescriptionItem struct {
	ID             uuid.UUID          `json:"id"`
	PrescriptionId uuid.UUID          `json:"prescription_id"`
	TimeZoneId     uuid.UUID          `json:"timezone_id"`
	MedicineId     uuid.UUID          `json:"medicine_id"`
	MedicineName   string             `json:"medicine_name"`
	TimeZoneName   string             `json:"timezone_name"`
	Midday         string             `json:"midday"`
	Hour           string             `json:"hour"`
	Minute         string             `json:"minute"`
	TakeAmount     float64            `json:"take_amount"`
	RemainAmount   float64            `json:"remain_amount"`
	TotalAmount    float64            `json:"total_amount"`
	MedicineUnit   string             `json:"medicine_unit"`
	Memo           string             `json:"memo"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
	Medicine       *medicine.Medicine `json:"medicine"`
}

func (p *Prescription) update(newModel *Prescription) {
	if newModel.HospitalName != "" {
		p.HospitalName = newModel.HospitalName
	}
	if newModel.PrescriptionName != "" {
		p.PrescriptionName = newModel.PrescriptionName
	}
	if newModel.TakeDays > 0 {
		p.TakeDays = newModel.TakeDays
	}
	if !newModel.StartedAt.IsZero() {
		p.StartedAt = newModel.StartedAt
	}
	if !newModel.FinishedAt.IsZero() {
		p.FinishedAt = newModel.FinishedAt
	}
	p.UpdatedAt = time.Now()
}

func (p *PrescriptionItem) update(newModel *PrescriptionItem) {
	if newModel.MedicineId != uuid.Nil {
		p.MedicineId = newModel.MedicineId
	}
	if newModel.MedicineName != "" {
		p.MedicineName = newModel.MedicineName
	}
	if newModel.TakeAmount > 0 {
		p.TakeAmount = newModel.TakeAmount
	}
	if newModel.MedicineUnit != "" {
		p.MedicineUnit = newModel.MedicineUnit
	}
	if newModel.Memo != "" {
		p.Memo = newModel.Memo
	}
	p.UpdatedAt = time.Now()
}
