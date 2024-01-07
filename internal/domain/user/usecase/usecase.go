package usecase

import (
	"nursing_api/internal/domain/user"
	"nursing_api/internal/domain/user/dto"
)

type UserUseCase interface {
	RegisterUser(req *dto.RegisterRequest) dto.RegisterResponse
	VerifyUser(user *user.User) (bool, error)
}
