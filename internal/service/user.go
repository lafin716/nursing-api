package service

import (
	"nursing_api/internal/common/utils"
	"nursing_api/internal/domain/user"
	"nursing_api/internal/domain/user/dto"
	"nursing_api/internal/domain/user/repository"
	"nursing_api/internal/domain/user/usecase"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) usecase.UserUseCase {
	return &userService{
		userRepo: userRepo,
	}
}

func (u userService) RegisterUser(req *dto.RegisterRequest) dto.RegisterResponse {
	count, err := u.userRepo.CountUserByEmail(req.Email)
	if err != nil {
		return dto.FailRegisterUser("회원 정보 조회 중 오류가 발생하였습니다", err)
	}
	if count > 0 {
		return dto.FailRegisterUser("이미 등록된 회원입니다.", nil)
	}

	newUser := &user.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: utils.GeneratePassword(req.Password),
	}
	savedUser, err := u.userRepo.CreateUser(newUser)
	if err != nil {
		return dto.FailRegisterUser("회원 정보 등록 중 오류가 발생하였습니다.", err)
	}

	return dto.OkRegisterUser(savedUser)
}

func (u userService) VerifyUser(user *user.User) (bool, error) {
	//TODO implement me
	panic("implement me")
}
