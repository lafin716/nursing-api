package prescription

import "time"

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
	prescription := &Prescription{
		UserId:           req.UserId,
		PrescriptionName: req.PrescriptionName,
		HospitalName:     req.HospitalName,
		TakeDays:         req.TakeDays,
		StartedAt:        req.StartedAt,
		FinishedAt:       req.FinishedAt,
		Memo:             req.Memo,
		CreatedAt:        time.Now(),
	}
	saved, err := p.repo.Add(prescription)
	if err != nil {
		return FailRegister("처방전 저장에 실패하였습니다.", err)
	}

	return OkRegister(saved.ID)
}
