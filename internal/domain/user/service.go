package user

import (
	"github.com/google/uuid"
	"log"
	"nursing_api/internal/common/utils"
	"strings"
)

type UseCase interface {
	RegisterUser(req *RegisterRequest) *RegisterResponse
	VerifyUser(req *LoginRequest) *LoginResponse
	GetUser(userId uuid.UUID) *GetUserResponse
	Leave(userId uuid.UUID) (bool, error)
	CheckDuplicatedEmail(req *CheckEmailRequest) *CheckEmailResponse
}

type service struct {
	userRepo Repository
}

func NewService(userRepo Repository) UseCase {
	return &service{
		userRepo: userRepo,
	}
}

func (u service) GetUser(userId uuid.UUID) *GetUserResponse {
	foundUser, err := u.userRepo.GetUserById(userId)
	if err != nil {
		return FailGetUserResponse("유저정보를 찾을 수 없습니다.", err)
	}

	return OkGetUserResponse(foundUser)
}

func (u service) RegisterUser(req *RegisterRequest) *RegisterResponse {
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

func (u service) VerifyUser(req *LoginRequest) *LoginResponse {
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

func (u service) Leave(userId uuid.UUID) (bool, error) {
	return u.userRepo.Delete(userId)
}

func (u service) CheckDuplicatedEmail(req *CheckEmailRequest) *CheckEmailResponse {
	foundUser, err := u.userRepo.GetUserByEmail(req.Email)
	if foundUser != nil {
		return FailCheckEmail("이미 등록된 이메일입니다.", nil)
	}

	if err != nil && !strings.Contains(err.Error(), "user not found") {
		log.Printf("err : %v", err)
		return FailCheckEmail("중복 확인 중 오류가 발생하였습니다.", err)
	}

	return OkCheckEmail()
}
