package user

import (
	"github.com/google/uuid"
	"nursing_api/internal/common/utils"
)

type userService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserUseCase {
	return &userService{
		userRepo: userRepo,
	}
}

func (u userService) GetUser(userId uuid.UUID) *GetUserResponse {
	foundUser, err := u.userRepo.GetUserById(userId)
	if err != nil {
		return FailGetUserResponse("유저정보를 찾을 수 없습니다.", err)
	}

	return OkGetUserResponse(foundUser)
}

func (u userService) RegisterUser(req *RegisterRequest) *RegisterResponse {
	count, err := u.userRepo.CountUserByEmail(req.Email)
	if err != nil {
		return FailRegisterUser("회원 정보 조회 중 오류가 발생하였습니다", err)
	}
	if count > 0 {
		return FailRegisterUser("이미 등록된 회원입니다.", nil)
	}

	newUser := &User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: utils.GeneratePassword(req.Password),
	}
	savedUser, err := u.userRepo.CreateUser(newUser)
	if err != nil {
		return FailRegisterUser("회원 정보 등록 중 오류가 발생하였습니다.", err)
	}

	return OkRegisterUser(savedUser)
}

func (u userService) VerifyUser(req *LoginRequest) *LoginResponse {
	foundUser, err := u.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return FailLoginUser("일치하는 정보를 찾을 수 없습니다.", err)
	}

	if !utils.ComparePassword(foundUser.PasswordHash, req.Password) {
		return FailLoginUser("일치하는 정보를 찾을 수 없습니다. 이메일 혹은 비밀번호를 다시 확인해주세요.", err)
	}

	foundUser.PasswordHash = ""
	return OkLoginUser(foundUser)
}

func (u userService) Leave(userId uuid.UUID) (bool, error) {
	return u.userRepo.Delete(userId)
}
