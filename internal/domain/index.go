package domain

import (
	"github.com/google/wire"
	"nursing_api/internal/domain/auth"
	"nursing_api/internal/domain/medicine"
	"nursing_api/internal/domain/plan"
	"nursing_api/internal/domain/prescription"
	"nursing_api/internal/domain/timezone"
	"nursing_api/internal/domain/user"
	medicine_api "nursing_api/pkg/api/medicine"
	"nursing_api/pkg/database"
	"nursing_api/pkg/jwt"
)

var Set = wire.NewSet(
	repository,
	service,
	gateway,
)

var repository = wire.NewSet(
	auth.NewRepository,
	user.NewRepository,
	medicine.NewRepository,
	prescription.NewRepository,
	timezone.NewRepository,
	plan.NewRepository,
)

var service = wire.NewSet(
	auth.NewService,
	user.NewService,
	medicine.NewService,
	prescription.NewService,
	plan.NewService,
	timezone.NewService,
)

var gateway = wire.NewSet(
	medicine_api.NewMedicineApi,
	jwt.NewJwtClient,
	database.NewPostgresClient,
)
