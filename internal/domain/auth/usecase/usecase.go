package usecase

import "nursing_api/internal/domain/auth/dto"

type AuthUseCase interface {
	SignIn(req *dto.SignInRequest) *dto.SignInResponse
	SignUp(req *dto.SignUpRequest) *dto.SignUpResponse
}
