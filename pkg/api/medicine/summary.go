package medicine

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

const (
	ENDPOINT = "http://apis.data.go.kr/1471000/DrbEasyDrugInfoService/getDrbEasyDrugList"
)

type SummaryApi interface {
	RetrievePill(pillName string) *SummaryResponse
}

type SummaryResponse struct {
	Header SummaryHeader `json:"header"`
	Body   SummaryBody   `json:"body"`
}

type SummaryHeader struct {
	ResultCode string `json:"resultCode"`
	ResultMsg  string `json:"resultMsg"`
}

type SummaryBody struct {
	PageNo     int           `json:"pageNo"`
	TotalCount int           `json:"totalCount"`
	NumOfRows  int           `json:"numOfRows"`
	Items      []SummaryItem `json:"items"`
}

type SummaryItem struct {
	EntpName            string      `json:"entpName"`
	ItemName            string      `json:"itemName"`
	ItemSeq             string      `json:"itemSeq"`
	EfcyQesitm          string      `json:"efcyQesitm"`
	UseMethodQesitm     string      `json:"useMethodQesitm"`
	AtpnWarnQesitm      interface{} `json:"atpnWarnQesitm"`
	AtpnQesitm          string      `json:"atpnQesitm"`
	IntrcQesitm         string      `json:"intrcQesitm"`
	SeQesitm            string      `json:"seQesitm"`
	DepositMethodQesitm string      `json:"depositMethodQesitm"`
	OpenDe              string      `json:"openDe"`
	UpdateDe            string      `json:"updateDe"`
	ItemImage           interface{} `json:"itemImage"`
	Bizrno              string      `json:"bizrno"`
}

type summaryApi struct {
	apiKey string
}

func NewSummaryApi(apiKey string) SummaryApi {
	return &summaryApi{
		apiKey: apiKey,
	}
}

func (s *summaryApi) RetrievePill(pillName string) *SummaryResponse {
	query := fmt.Sprintf("?type=json&serviceKey=%s&itemName=%s", s.apiKey, pillName)
	agent := fiber.Get(ENDPOINT + query)

	_, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return nil
	}

	var response SummaryResponse
	err := json.Unmarshal(body, &response)
	if err != nil {
		return nil
	}

	return &response
}

func getAgent(method string) *fiber.Agent {
	agent := fiber.AcquireAgent()
	agent.ContentType("application/json")
	agent.Request().Header.SetMethod(method)
	agent.Request().SetRequestURI(ENDPOINT)

	return agent
}
