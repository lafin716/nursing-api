package persistence

import "github.com/google/wire"

var Set = wire.NewSet(
	NewUserRepository,
	NewAuthRepository,
)
