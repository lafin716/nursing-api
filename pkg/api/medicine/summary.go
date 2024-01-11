package medicine_api

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

const (
	SUMMARY_ENDPOINT = "http://apis.data.go.kr/1471000/DrbEasyDrugInfoService/getDrbEasyDrugList"
)

type SummaryApi interface {
	RetrievePill(pillName string) *Response[SummaryItem]
}

type SummaryItem struct {
	EntpName            string `json:"entpName"`
	ItemName            string `json:"itemName"`
	ItemSeq             string `json:"itemSeq"`
	EfcyQesitm          string `json:"efcyQesitm"`
	UseMethodQesitm     string `json:"useMethodQesitm"`
	AtpnWarnQesitm      string `json:"atpnWarnQesitm"`
	AtpnQesitm          string `json:"atpnQesitm"`
	IntrcQesitm         string `json:"intrcQesitm"`
	SeQesitm            string `json:"seQesitm"`
	DepositMethodQesitm string `json:"depositMethodQesitm"`
	OpenDe              string `json:"openDe"`
	UpdateDe            string `json:"updateDe"`
	ItemImage           string `json:"itemImage"`
	Bizrno              string `json:"bizrno"`
}

type summaryApi struct {
	apiKey string
}

func NewSummaryApi(apiKey string) SummaryApi {
	return &summaryApi{
		apiKey: apiKey,
	}
}

func (s *summaryApi) RetrievePill(pillName string) *Response[SummaryItem] {
	query := fmt.Sprintf("?type=json&serviceKey=%s&itemName=%s", s.apiKey, pillName)
	agent := fiber.Get(SUMMARY_ENDPOINT + query)

	_, body, errs := agent.Bytes()
	if len(errs) > 0 {
		log.Println(errs)
		return nil
	}

	var response Response[SummaryItem]
	err := json.Unmarshal(body, &response)
	if err != nil {
		return nil
	}

	return &response
}
