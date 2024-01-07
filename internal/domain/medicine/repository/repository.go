package medicine

import "nursing_api/internal/domain/medicine"

type MedicineRepository interface {
	SavePill(medicine medicine.Medicine) (bool, error)
}
