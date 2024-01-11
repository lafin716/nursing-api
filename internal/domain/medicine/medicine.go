package medicine

import (
	"time"
)

type Medicine struct {
	Name         string
	ItemSeq      string
	CompanyName  string
	Description  string
	Usage        string
	Effect       string
	SideEffect   string
	Caution      string
	Warning      string
	Interaction  string
	KeepMethod   string
	Appearance   string
	ColorClass1  string
	ColorClass2  string
	PillImage    string
	ClassName    string
	OtcName      string
	FormCodeName string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type MedicineRepository interface {
	SavePills(medicine []*Medicine) (int, error)
	SavePill(medicine *Medicine) (bool, error)
	GetPillsByNames(pillName string) ([]*Medicine, error)
}

type MedicineUseCase interface {
	SearchMedicine(pillName string) *SearchPillResponse
}
