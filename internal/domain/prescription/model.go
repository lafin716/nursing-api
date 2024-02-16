package prescription

import (
	"github.com/google/uuid"
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
	Memo              string              `json:"memo"`
	CreatedAt         time.Time           `json:"created_at"`
	UpdatedAt         time.Time           `json:"updated_at"`
	PrescriptionItems []*PrescriptionItem `json:"prescription_items"`
}

type PrescriptionItem struct {
	ID             uuid.UUID `json:"id"`
	TimeZoneLinkId uuid.UUID `json:"timezone_link_id"`
	MedicineId     uuid.UUID `json:"medicine_id"`
	MedicineName   string    `json:"medicine_name"`
	TakeAmount     float64   `json:"take_amount"`
	MedicineUnit   string    `json:"medicine_unit"`
	Memo           string    `json:"memo"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Repository interface {
	GetById(id uuid.UUID) (*Prescription, error)
	GetListByUserId(search *PrescriptionSearch) ([]*Prescription, error)
	GetItemListByPrescriptionId(timezoneId uuid.UUID) ([]*PrescriptionItem, error)
	GetItemListBySearch(search *PrescriptionSearch) ([]*PrescriptionItem, error)
	Add(prescription *Prescription) (*Prescription, error)
	Update(prescription *Prescription) (int, error)
	Delete(id uuid.UUID) (bool, error)
	GetItemById(itemId uuid.UUID) (*PrescriptionItem, error)
	AddItem(prescriptionId uuid.UUID, item *PrescriptionItem) (*PrescriptionItem, error)
	UpdateItem(prescriptionItem *PrescriptionItem) (int, error)
	DeleteItem(itemId uuid.UUID) (bool, error)
}

type UseCase interface {
	GetByDate(req *GetByDateRequest) *GetByDateResponse
	GetList(req *GetListRequest) *GetListResponse
	Register(req *RegisterRequest) *RegisterResponse
	Update(req *UpdateRequest) *UpdateResponse
	Delete(req *DeleteRequest) *DeleteResponse
	AddItem(req *AddItemRequest) *AddItemResponse
	UpdateItem(req *UpdateItemRequest) *UpdateItemResponse
	DeleteItem(req *DeleteItemRequest) *DeleteItemResponse
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
	if newModel.Memo != "" {
		p.Memo = newModel.Memo
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
