package medicine

import (
	"net/url"
	"nursing_api/internal/common/dto"
	"nursing_api/internal/common/response"
	medicine_api "nursing_api/pkg/api/medicine"
	"strings"
)

type UseCase interface {
	Search(pillName string) dto.BaseResponse[[]*Medicine]
}

type medicineService struct {
	medicineRepo Repository
	medicineApi  medicine_api.MedicineApi
}

func NewService(
	medicineRepo Repository,
	medicineApi medicine_api.MedicineApi,
) UseCase {
	return &medicineService{
		medicineRepo: medicineRepo,
		medicineApi:  medicineApi,
	}
}

func (m medicineService) Search(pillName string) dto.BaseResponse[[]*Medicine] {
	// DB에 이미 존재하는 경우 API 통신하지않고 바로 반환
	decoded, err := url.QueryUnescape(pillName)
	decoded = strings.ReplaceAll(decoded, "\b", "")
	if err != nil {
		return dto.Fail[[]*Medicine](response.CODE_NOT_AVAILABLE_MEDICINE_KEY, err)
	}

	medicines, _ := m.medicineRepo.GetPillsByNames(decoded)
	if medicines != nil {
		return dto.Ok[[]*Medicine](response.CODE_SUCCESS, &medicines)
	}

	// 의약품 API를 통해 가져온다
	resp := m.medicineApi.Summary().RetrievePill(pillName)
	if resp == nil {
		return dto.Fail[[]*Medicine](response.CODE_NOT_FOUND_MEDICINE, err)
	}
	summaryPills := m.fetchSummary(pillName)
	if summaryPills == nil || len(summaryPills) == 0 {
		return dto.Fail[[]*Medicine](response.CODE_SEARCH_MEDICINE_ZERO, err)
	}

	// API 결과를 DB에 저장
	okCount, err := m.medicineRepo.SavePills(summaryPills)
	if err != nil {
		return dto.Fail[[]*Medicine](response.CODE_FAIL_SAVE_MEDICINE_DATA, err)
	}
	if okCount == 0 {
		return dto.Fail[[]*Medicine](response.CODE_FAIL_SAVE_MEDICINE_ZERO, nil)
	}

	// 저장된 데이터를 다시 조회하여 반환
	medicines, _ = m.medicineRepo.GetPillsByNames(decoded)
	if medicines != nil {
		return dto.Ok[[]*Medicine](response.CODE_SUCCESS, &medicines)
	} else {
		// 저장된 데이터가 없을 경우 API 결과를 반환
		apiLimit := 10
		if len(summaryPills) < apiLimit {
			apiLimit = len(summaryPills)
		}
		limitedPills := summaryPills[:apiLimit]
		return dto.Ok[[]*Medicine](response.CODE_SUCCESS, &limitedPills)
	}
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
