package user

import (
	"github.com/google/uuid"
	"log"
	"nursing_api/internal/common/dto"
	"nursing_api/internal/common/response"
	"nursing_api/internal/common/utils"
	"strings"
)

type UseCase interface {
	GetUser(userId uuid.UUID) dto.BaseResponse[User]
	RegisterUser(req *RegisterRequest) dto.BaseResponse[User]
	VerifyUser(req *LoginRequest) dto.BaseResponse[User]
	Leave(userId uuid.UUID) (bool, error)
	CheckDuplicatedEmail(req *CheckEmailRequest) dto.BaseResponse[any]
}

type service struct {
	userRepo Repository
}

func NewService(userRepo Repository) UseCase {
	return &service{
		userRepo: userRepo,
	}
}

func (u service) GetUser(userId uuid.UUID) dto.BaseResponse[User] {
	foundUser, err := u.userRepo.GetUserById(userId)
	if err != nil {
		return dto.Fail[User](response.CODE_USER_NOTFOUND, err)
	}

	return dto.Ok[User](response.CODE_SUCCESS, foundUser)
}

func (u service) RegisterUser(req *RegisterRequest) dto.BaseResponse[User] {
	count, err := u.userRepo.CountUserByEmail(req.Email)
	if err != nil {
		return dto.Fail[User](response.CODE_OCCUR_ERROR_DURING_GET_USER, err)
	}
	if count > 0 {
		return dto.Fail[User](response.CODE_DUPLICATE_EMAIL, err)
	}

	newUser := &User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: utils.GeneratePassword(req.Password),
	}
	savedUser, err := u.userRepo.CreateUser(newUser)
	if err != nil {
		return dto.Fail[User](response.CODE_OCCUR_ERROR_DURING_SAVE_USER, err)
	}

	return dto.Ok[User](response.CODE_SUCCESS, savedUser)
}

func (u service) VerifyUser(req *LoginRequest) dto.BaseResponse[User] {
	foundUser, err := u.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return dto.Fail[User](response.CODE_USER_NOTFOUND, err)
	}

	if !utils.ComparePassword(foundUser.PasswordHash, req.Password) {
		return dto.Fail[User](response.CODE_NOT_MATCHING_SIGN_IN_INFORMATION, err)
	}

	foundUser.PasswordHash = ""
	return dto.Ok[User](response.CODE_SUCCESS, foundUser)
}

func (u service) Leave(userId uuid.UUID) (bool, error) {
	return u.userRepo.Delete(userId)
}

func (u service) CheckDuplicatedEmail(req *CheckEmailRequest) dto.BaseResponse[any] {
	foundUser, err := u.userRepo.GetUserByEmail(req.Email)
	if foundUser != nil {
		return dto.Fail[any](response.CODE_DUPLICATE_EMAIL, err)
	}

	if err != nil && !strings.Contains(err.Error(), "user not found") {
		log.Printf("err : %v", err)
		return dto.Fail[any](response.CODE_OCCUR_ERROR_DURING_CHECK_EMAIL, err)
	}

	return dto.Ok[any](response.CODE_AVAILABLE_EMAIL, nil)
}
