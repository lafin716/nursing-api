package medicine

import "github.com/google/wire"

var Set = wire.NewSet(
	NewMedicineApi,
)
