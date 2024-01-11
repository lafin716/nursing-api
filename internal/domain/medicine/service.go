package medicine

import (
	"net/url"
	medicine_api "nursing_api/pkg/api/medicine"
	"strings"
)

type medicineService struct {
	medicineRepo MedicineRepository
	medicineApi  medicine_api.MedicineApi
}

func NewMedicineService(medicineRepo MedicineRepository, medicineApi medicine_api.MedicineApi) MedicineUseCase {
	return &medicineService{
		medicineRepo: medicineRepo,
		medicineApi:  medicineApi,
	}
}

func (m medicineService) SearchMedicine(pillName string) *SearchPillResponse {
	// DB에 이미 존재하는 경우 API 통신하지않고 바로 반환
	decoded, err := url.QueryUnescape(pillName)
	decoded = strings.ReplaceAll(decoded, "\b", "")
	if err != nil {
		return FailSearchPill("검색어가 유효하지 않습니다.", err)
	}

	medicines, _ := m.medicineRepo.GetPillsByNames(decoded)
	if medicines != nil {
		return OkSearchPill(medicines)
	}

	// 의약품 API를 통해 가져온다
	resp := m.medicineApi.Summary().RetrievePill(pillName)
	if resp == nil {
		return FailSearchPill("의약품 정보를 찾을 수 없습니다.", nil)
	}
	summaryPills := m.fetchSummary(pillName)
	if summaryPills == nil || len(summaryPills) == 0 {
		return FailSearchPill("의약품 검색 결과: 0건", nil)
	}

	// API 결과를 DB에 저장
	okCount, err := m.medicineRepo.SavePills(summaryPills)
	if err != nil {
		return FailSearchPill("의약품 검색 결과 저장 실패", err)
	}
	if okCount == 0 {
		return FailSearchPill("의약품 검색 결과 저장: 0건", nil)
	}

	apiLimit := 10
	if len(summaryPills) < apiLimit {
		apiLimit = len(summaryPills)
	}

	return OkSearchPill(summaryPills[:apiLimit])
}

// 의약품 개요정보 조회
func (m medicineService) fetchSummary(pillName string) []*Medicine {
	summary := m.medicineApi.Summary().RetrievePill(pillName)
	if summary == nil {
		return nil
	}

	appearance := m.medicineApi.Appearance().RetrievePill(pillName)
	if appearance == nil {
		return nil
	}

	pillSet := make(map[string]*Medicine)
	for _, item := range summary.Body.Items {
		trimPillName := strings.TrimSpace(item.ItemName)
		pillSet[item.ItemSeq] = &Medicine{
			Name:        trimPillName,
			ItemSeq:     item.ItemSeq,
			Usage:       item.UseMethodQesitm,
			Effect:      item.EfcyQesitm,
			SideEffect:  item.SeQesitm,
			Caution:     item.AtpnQesitm,
			Warning:     item.AtpnWarnQesitm,
			Interaction: item.IntrcQesitm,
			PillImage:   item.ItemImage,
			KeepMethod:  item.DepositMethodQesitm,
		}
	}

	for _, item := range appearance.Body.Items {
		trimPillName := strings.TrimSpace(item.ItemName)
		if pillSet[item.ItemSeq] == nil {
			break
		}

		pillSet[item.ItemSeq].Name = trimPillName
		pillSet[item.ItemSeq].ItemSeq = item.ItemSeq
		pillSet[item.ItemSeq].Appearance = item.DrugShape
		pillSet[item.ItemSeq].Description = item.Chart
		pillSet[item.ItemSeq].ColorClass1 = item.ColorClass1
		pillSet[item.ItemSeq].ColorClass2 = item.ColorClass2
		pillSet[item.ItemSeq].PillImage = item.ItemImage
		pillSet[item.ItemSeq].OtcName = item.EtcOtcName
		pillSet[item.ItemSeq].FormCodeName = item.FormCodeName
	}

	var medicines []*Medicine
	for _, item := range pillSet {
		medicines = append(medicines, item)
	}

	return medicines
}
