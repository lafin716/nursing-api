package prescription

import (
	"fmt"
	"nursing_api/pkg/jwt"
	"time"
)

const (
	DATE_LAYOUT = "2006-01-02"
)

type prescriptionService struct {
	repo      PrescriptionRepository
	jwtClient *jwt.JwtClient
}

func NewPrescriptionService(
	repo PrescriptionRepository,
	jwtClient *jwt.JwtClient,
) PrescriptionUseCase {
	return &prescriptionService{
		repo:      repo,
		jwtClient: jwtClient,
	}
}

func (p prescriptionService) GetList(req *GetListRequest) *GetListResponse {
	defaultFormat := "2006-01-02"
	targetDate, err := time.Parse(defaultFormat, req.Date)
	if err != nil {
		targetDate, _ = time.Parse(defaultFormat, time.Now().Format(defaultFormat))
	}
	if req.Limit == 0 {
		req.Limit = 10
	}
	search := &PrescriptionSearch{
		UserId:     req.UserId,
		TargetDate: targetDate,
		Limit:      req.Limit,
	}
	resp, err := p.repo.GetListByUserId(search)
	for _, item := range resp {
		prescriptionItems, err := p.repo.GetItemListByPrescriptionId(item.ID)
		if err != nil {
			item.PrescriptionItems = []*PrescriptionItem{}
			continue
		}
		item.PrescriptionItems = prescriptionItems
	}
	if err != nil {
		return FailGetList("처방전 목록이 없습니다.", err)
	}
	if len(resp) == 0 {
		return OkGetListMessage("등록된 처방전이 없습니다.")
	}

	return OkGetList(resp)
}

// TODO 처방전 날짜 다시 협의 필요 복용일, 복용시작, 복용종료
func (p prescriptionService) Register(req *RegisterRequest) *RegisterResponse {
	started, err := time.Parse(DATE_LAYOUT, req.StartedAt)
	if err != nil {
		return FailRegister("복용시작 날짜형식이 맞지않습니다.", err)
	}

	finished, err := time.Parse(DATE_LAYOUT, req.FinishedAt)
	if err != nil {
		return FailRegister("복용종료 날짜형식이 맞지않습니다.", err)
	}

	fmt.Printf("처방전 등록 파라미터: %+v", req)

	items := []*PrescriptionItem{}
	for _, item := range req.Items {
		items = append(items, &PrescriptionItem{
			UserId:       req.UserId,
			MedicineId:   item.MedicineId,
			MedicineName: item.MedicineName,
			TakeTimeZone: item.TakeTimeZone,
			TakeMoment:   item.TakeMoment,
			TakeEtc:      item.TakeEtc,
			TakeAmount:   item.TakeAmount,
			MedicineUnit: item.MedicineUnit,
			CreatedAt:    time.Now(),
		})
	}

	prescription := &Prescription{
		UserId:            req.UserId,
		PrescriptionName:  req.PrescriptionName,
		HospitalName:      req.HospitalName,
		TakeDays:          req.TakeDays,
		StartedAt:         started,
		FinishedAt:        finished,
		Memo:              req.Memo,
		CreatedAt:         time.Now(),
		PrescriptionItems: items,
	}
	saved, err := p.repo.Add(prescription)
	if err != nil {
		return FailRegister("처방전 저장에 실패하였습니다.", err)
	}

	return OkRegister(saved.ID)
}

func (p prescriptionService) Update(req *UpdateRequest) *UpdateResponse {
	found, err := p.repo.GetById(req.ID)
	if err != nil {
		return FailUpdate("처방전 데이터를 찾을 수 없습니다.", err)
	}

	started := found.StartedAt
	if req.StartedAt != "" {
		started, err = time.Parse(DATE_LAYOUT, req.StartedAt)
		if err != nil {
			return FailUpdate("복용시작 날짜형식이 맞지않습니다.", err)
		}
	}

	finished := found.FinishedAt
	if req.FinishedAt != "" {
		finished, err = time.Parse(DATE_LAYOUT, req.FinishedAt)
		if err != nil {
			return FailUpdate("복용종료 날짜형식이 맞지않습니다.", err)
		}
	}

	updated := &Prescription{
		ID:               req.ID,
		PrescriptionName: req.PrescriptionName,
		HospitalName:     req.HospitalName,
		TakeDays:         req.TakeDays,
		StartedAt:        started,
		FinishedAt:       finished,
		Memo:             req.Memo,
	}
	found.update(updated)

	result, err := p.repo.Update(found)
	if err != nil {
		return FailUpdate("처방전 업데이트 중 오류가 발생하였습니다.", err)
	}
	if result <= 0 {
		return FailUpdate("처방전이 업데이트 되지않았습니다.", nil)
	}

	return OkUpdate(found)
}

func (p prescriptionService) Delete(req *DeleteRequest) *DeleteResponse {
	found, err := p.repo.GetById(req.ID)
	if err != nil || found == nil {
		return FailDelete("처방전 데이터를 찾을 수 없습니다.", nil)
	}

	result, err := p.repo.Delete(req.ID)
	if err != nil {
		return FailDelete("처방전 삭제가 실패하였습니다.", err)
	}

	if !result {
		return FailDelete("처방전 삭제가 실패하였습니다.", err)
	}

	return OkDelete()
}

func (p prescriptionService) AddItem(req *AddItemRequest) *AddItemResponse {
	found, err := p.repo.GetById(req.PrescriptionId)
	if err != nil {
		return FailAddItem("데이터를 찾을 수 없습니다.", err)
	}

	newItem := &PrescriptionItem{
		UserId:         found.UserId,
		PrescriptionId: found.ID,
		MedicineId:     req.MedicineId,
		MedicineName:   req.MedicineName,
		TakeTimeZone:   req.TakeTimeZone,
		TakeMoment:     req.TakeMoment,
		TakeEtc:        req.TakeEtc,
		TakeAmount:     req.TakeAmount,
		MedicineUnit:   req.MedicineUnit,
		Memo:           req.Memo,
		CreatedAt:      time.Now(),
	}

	_, err = p.repo.AddItem(found.ID, newItem)
	if err != nil {
		return FailAddItem("처방전 의약품 추가 중 오류가 발생하였습니다.", err)
	}

	return OkAddItem()
}

func (p prescriptionService) UpdateItem(req *UpdateItemRequest) *UpdateItemResponse {
	found, err := p.repo.GetItemById(req.ID)
	if err != nil {
		return FailUpdateItem("처방전 의약품 데이터를 찾을 수 없습니다.", err)
	}

	newItem := &PrescriptionItem{
		ID:           req.ID,
		MedicineId:   req.MedicineId,
		MedicineName: req.MedicineName,
		TakeTimeZone: req.TakeTimeZone,
		TakeMoment:   req.TakeMoment,
		TakeEtc:      req.TakeEtc,
		TakeAmount:   req.TakeAmount,
		MedicineUnit: req.MedicineUnit,
		Memo:         req.Memo,
	}
	found.update(newItem)
	_, err = p.repo.UpdateItem(found)
	if err != nil {
		return FailUpdateItem("처방전 의약품 업데이트 중 오류가 발생하였습니다.", err)
	}

	return OkUpdateItem()
}

func (p prescriptionService) DeleteItem(req *DeleteItemRequest) *DeleteItemResponse {
	found, err := p.repo.GetItemById(req.ID)
	if err != nil {
		return FailDeleteItem("처방전 의약품 데이터를 찾을 수 없습니다.", err)
	}

	_, err = p.repo.DeleteItem(found.ID)
	if err != nil {
		return FailDeleteItem("처방전 의약품 삭제 중 오류가 발생하였습니다.", err)
	}

	return OkDeleteItem()
}

func (p prescriptionService) GetByDate(req *GetByDateRequest) *GetByDateResponse {
	//TODO implement me
	panic("implement me")
}
