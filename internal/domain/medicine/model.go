package medicine

import (
	"github.com/google/uuid"
	"time"
)

type Medicine struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	ItemSeq      string    `json:"item_seq"`
	CompanyName  string    `json:"company_name"`
	Description  string    `json:"description"`
	Usage        string    `json:"usage"`
	Effect       string    `json:"effect"`
	SideEffect   string    `json:"side_effect"`
	Caution      string    `json:"caution"`
	Warning      string    `json:"warning"`
	Interaction  string    `json:"interaction"`
	KeepMethod   string    `json:"keep_method"`
	Appearance   string    `json:"appearance"`
	ColorClass1  string    `json:"color_class1"`
	ColorClass2  string    `json:"color_class2"`
	PillImage    string    `json:"pill_image"`
	ClassName    string    `json:"class_name"`
	OtcName      string    `json:"otc_name"`
	FormCodeName string    `json:"form_code_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
