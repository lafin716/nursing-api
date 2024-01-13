package prescription

import "time"

const (
	DATE_LAYOUT = "2006-01-02"
)

type prescriptionService struct {
	repo PrescriptionRepository
}

func NewPrescriptionService(
	repo PrescriptionRepository,
) PrescriptionUseCase {
	return &prescriptionService{
		repo: repo,
	}
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

	prescription := &Prescription{
		UserId:           req.UserId,
		PrescriptionName: req.PrescriptionName,
		HospitalName:     req.HospitalName,
		TakeDays:         req.TakeDays,
		StartedAt:        started,
		FinishedAt:       finished,
		Memo:             req.Memo,
		CreatedAt:        time.Now(),
	}
	saved, err := p.repo.Add(prescription)
	if err != nil {
		return FailRegister("처방전 저장에 실패하였습니다.", err)
	}

	return OkRegister(saved.ID)
}
