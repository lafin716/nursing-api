package usecase

import (
	"github.com/google/uuid"
	"nursing_api/internal/domain/user/dto"
)

type UserUseCase interface {
	RegisterUser(req *dto.RegisterRequest) *dto.RegisterResponse
	VerifyUser(req *dto.LoginRequest) *dto.LoginResponse
	GetUser(userId uuid.UUID) *dto.GetUserResponse
}
