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
	resp, err := p.repo.GetListByUserId(req.UserId)
	if err != nil {
		return FailGetList("처방전 목록이 없습니다.", err)
	}

	return OkGetList(resp)
}

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
