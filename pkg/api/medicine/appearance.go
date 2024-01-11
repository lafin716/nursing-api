package medicine_api

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

const (
	APPEARANCE_ENDPOINT = "http://apis.data.go.kr/1471000/MdcinGrnIdntfcInfoService01/getMdcinGrnIdntfcInfoList01"
)

type AppearanceApi interface {
	RetrievePill(pillName string) *Response[AppearanceItem]
}

type AppearanceItem struct {
	ItemSeq           string      `json:"ITEM_SEQ"`
	ItemName          string      `json:"ITEM_NAME"`
	EntpSeq           string      `json:"ENTP_SEQ"`
	EntpName          string      `json:"ENTP_NAME"`
	Chart             string      `json:"CHART"`
	ItemImage         string      `json:"ITEM_IMAGE"`
	PrintFront        string      `json:"PRINT_FRONT"`
	PrintBack         interface{} `json:"PRINT_BACK"`
	DrugShape         string      `json:"DRUG_SHAPE"`
	ColorClass1       string      `json:"COLOR_CLASS1"`
	ColorClass2       string      `json:"COLOR_CLASS2"`
	LineFront         interface{} `json:"LINE_FRONT"`
	LineBack          interface{} `json:"LINE_BACK"`
	LengLong          string      `json:"LENG_LONG"`
	LengShort         string      `json:"LENG_SHORT"`
	Thick             string      `json:"THICK"`
	ImgRegistTs       string      `json:"IMG_REGIST_TS"`
	ClassNo           string      `json:"CLASS_NO"`
	ClassName         string      `json:"CLASS_NAME"`
	EtcOtcName        string      `json:"ETC_OTC_NAME"`
	ItemPermitDate    string      `json:"ITEM_PERMIT_DATE"`
	FormCodeName      string      `json:"FORM_CODE_NAME"`
	MarkCodeFrontAnal string      `json:"MARK_CODE_FRONT_ANAL"`
	MarkCodeBackAnal  string      `json:"MARK_CODE_BACK_ANAL"`
	MarkCodeFrontImg  string      `json:"MARK_CODE_FRONT_IMG"`
	MarkCodeBackImg   string      `json:"MARK_CODE_BACK_IMG"`
	ItemEngName       interface{} `json:"ITEM_ENG_NAME"`
	ChangeDate        interface{} `json:"CHANGE_DATE"`
	MarkCodeFront     interface{} `json:"MARK_CODE_FRONT"`
	MarkCodeBack      interface{} `json:"MARK_CODE_BACK"`
	EdiCode           interface{} `json:"EDI_CODE"`
	Bizrno            string      `json:"BIZRNO"`
}

type appearanceApi struct {
	apiKey string
}

func NewAppearanceApi(apiKey string) AppearanceApi {
	return &appearanceApi{
		apiKey: apiKey,
	}
}

func (s *appearanceApi) RetrievePill(pillName string) *Response[AppearanceItem] {
	query := fmt.Sprintf("?type=json&serviceKey=%s&item_name=%s", s.apiKey, pillName)
	agent := fiber.Get(APPEARANCE_ENDPOINT + query)

	_, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return nil
	}

	var response Response[AppearanceItem]
	err := json.Unmarshal(body, &response)
	if err != nil {
		return nil
	}

	return &response
}
