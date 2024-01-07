package service

import (
	"github.com/google/uuid"
	"nursing_api/internal/common/utils"
	"nursing_api/internal/domain/user"
	"nursing_api/internal/domain/user/dto"
	"nursing_api/internal/domain/user/repository"
	"nursing_api/internal/domain/user/usecase"
)

type userService struct {
	userRepo repository.UserRepository
}

func (u userService) GetUser(userId uuid.UUID) *dto.GetUserResponse {
	foundUser, err := u.userRepo.GetUserById(userId)
	if err != nil {
		return dto.FailGetUserResponse("유저정보를 찾을 수 없습니다.", err)
	}

	return dto.OkGetUserResponse(foundUser)
}

func NewUserService(userRepo repository.UserRepository) usecase.UserUseCase {
	return &userService{
		userRepo: userRepo,
	}
}

func (u userService) RegisterUser(req *dto.RegisterRequest) *dto.RegisterResponse {
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

func (u userService) VerifyUser(req *dto.LoginRequest) *dto.LoginResponse {
	foundUser, err := u.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return dto.FailLoginUser("일치하는 정보를 찾을 수 없습니다.", err)
	}

	if !utils.ComparePassword(foundUser.PasswordHash, req.Password) {
		return dto.FailLoginUser("일치하는 정보를 찾을 수 없습니다. 이메일 혹은 비밀번호를 다시 확인해주세요.", err)
	}

	foundUser.PasswordHash = ""
	return dto.OkLoginUser(foundUser)
}
