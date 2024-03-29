package config

import "github.com/google/wire"

var Set = wire.NewSet(
	NewFiberConfig,
	NewJwtConfig,
	NewDatabaseConfig,
	NewMedicineApiConfig,
)
