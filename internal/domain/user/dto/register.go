package dto

import (
	"nursing_api/internal/domain/user"
)

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,lte=100"`
	Email    string `json:"email" validate:"required,email,lte=100"`
	Password string `json:"password" validate:"required,lte=100"`
}

type RegisterResponse struct {
	Success bool
	Message string
	User    *user.User
	Error   error
}

func OkRegisterUser(newUser *user.User) *RegisterResponse {
	return &RegisterResponse{
		Success: true,
		Message: "",
		User:    newUser,
	}
}

func FailRegisterUser(message string, err error) *RegisterResponse {
	return &RegisterResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
