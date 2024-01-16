package prescription

import (
	"github.com/google/uuid"
	"time"
)

type Prescription struct {
	ID                uuid.UUID           `json:"id"`
	UserId            uuid.UUID           `json:"user_id"`
	PrescriptionName  string              `json:"prescription_name"`
	HospitalName      string              `json:"hospital_name"`
	TakeDays          int                 `json:"take_days"`
	StartedAt         time.Time           `json:"started_at"`
	FinishedAt        time.Time           `json:"finished_at"`
	Memo              string              `json:"memo"`
	CreatedAt         time.Time           `json:"created_at"`
	UpdatedAt         time.Time           `json:"updated_at"`
	PrescriptionItems []*PrescriptionItem `json:"prescription_items"`
}

type PrescriptionItem struct {
	ID             uuid.UUID `json:"id"`
	UserId         uuid.UUID `json:"user_id"`
	PrescriptionId uuid.UUID `json:"prescription_id"`
	MedicineName   string    `json:"medicine_name"`
	TakeTimeZone   string    `json:"take_time_zone"`
	TakeMoment     string    `json:"take_moment"`
	TakeEtc        string    `json:"take_etc"`
	TakeAmount     float64   `json:"take_amount"`
	MedicineUnit   string    `json:"medicine_unit"`
	Memo           string    `json:"memo"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type PrescriptionRepository interface {
	GetById(id uuid.UUID) (*Prescription, error)
	GetListByUserId(search *PrescriptionSearch) ([]*Prescription, error)
	GetItemListByPrescriptionId(prescriptionId uuid.UUID) ([]*PrescriptionItem, error)
	Add(prescription *Prescription) (*Prescription, error)
	AddItem(item PrescriptionItem) (*PrescriptionItem, error)
	Update(prescription *Prescription) (int, error)
	UpdateItem(itemId uuid.UUID) (int, error)
	Delete(id uuid.UUID) (bool, error)
	DeleteItem(itemId uuid.UUID) (bool, error)
}

type PrescriptionUseCase interface {
	GetList(req *GetListRequest) *GetListResponse
	Register(req *RegisterRequest) *RegisterResponse
	Update()
	Remove()
	AddItem()
	UpdateItem()
	DeleteItem()
}
