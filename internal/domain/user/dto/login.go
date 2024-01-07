package dto

import "nursing_api/internal/domain/user"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,lte=100"`
	Password string `json:"password" validate:"required,lte=100"`
}

type LoginResponse struct {
	Success bool
	Message string
	User    *user.User
	Error   error
}

func OkLoginUser(loginUser *user.User) *LoginResponse {
	return &LoginResponse{
		Success: true,
		Message: "",
		User:    loginUser,
	}
}

func FailLoginUser(message string, err error) *LoginResponse {
	return &LoginResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
