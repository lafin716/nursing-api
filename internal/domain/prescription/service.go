package prescription

import (
	"nursing_api/pkg/jwt"
	"nursing_api/pkg/mono"
	"strings"
	"time"
)

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

type prescriptionService struct {
	repo      Repository
	mono      *mono.Client
	jwtClient *jwt.JwtClient
}

func NewService(
	repo Repository,
	jwtClient *jwt.JwtClient,
) UseCase {
	return &prescriptionService{
		repo:      repo,
		jwtClient: jwtClient,
	}
}

func (p prescriptionService) GetList(req *GetListRequest) *GetListResponse {
	targetDate, err := p.mono.Date.Parse("Y-m-d", req.Date)
	if err != nil {
		targetDate = time.Now()
	}
	if req.Limit == 0 {
		req.Limit = 10
	}
	search := &SearchCondition{
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

func (p prescriptionService) Register(req *RegisterRequest) *RegisterResponse {
	started, err := p.mono.Date.Parse("Y-m-d", req.StartedAt)
	if err != nil {
		return FailRegister("복용시작 날짜형식이 맞지않습니다.", err)
	}

	// 복용 종료일 계산
	var finished time.Time
	if strings.TrimSpace(req.FinishedAt) != "" {
		// 복용 종료일이 파라미터에 있는 경우 그대로 사용
		finished, err = p.mono.Date.Parse("Y-m-d", req.FinishedAt)
		if err != nil {
			return FailRegister("복용종료 날짜형식이 맞지않습니다.", err)
		}
	} else {
		// 복용종료일시가 없는 경우 복용시작일시를 기준으로 복용일을 더한다
		finished = started.AddDate(0, 0, 1*req.TakeDays)
	}

	items := []*PrescriptionItem{}
	for _, item := range req.Items {
		items = append(items, &PrescriptionItem{
			MedicineId:   item.MedicineId,
			MedicineName: item.MedicineName,
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
		started, err = p.mono.Date.Parse("Y-m-d", req.StartedAt)
		if err != nil {
			return FailUpdate("복용시작 날짜형식이 맞지않습니다.", err)
		}
	}

	finished := found.FinishedAt
	if req.FinishedAt != "" {
		finished, err = p.mono.Date.Parse("Y-m-d", req.FinishedAt)
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
