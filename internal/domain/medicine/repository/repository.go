package medicine

import "nursing_api/internal/domain/medicine"

type MedicineRepository interface {
	SavePills(medicine medicine.Medicine) (bool, error)
}
