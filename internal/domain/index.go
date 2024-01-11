package domain

import (
	"github.com/google/wire"
	"nursing_api/internal/domain/auth"
	"nursing_api/internal/domain/medicine"
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
	auth.NewAuthRepository,
	user.NewUserRepository,
)

var service = wire.NewSet(
	auth.NewAuthService,
	user.NewUserService,
	medicine.NewMedicineService,
)

var gateway = wire.NewSet(
	medicine_api.NewMedicineApi,
	jwt.NewJwtClient,
	database.NewPostgresClient,
)
