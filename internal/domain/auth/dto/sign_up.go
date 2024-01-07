package dto

import (
	"nursing_api/internal/domain/auth"
)

type SignUpRequest struct {
	Name     string `json:"name" validate:"required,lte=100"`
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}

type SignUpResponse struct {
	Success bool
	Message string
	Token   *auth.Token
	Error   error
}

func OkSignUp(token *auth.Token) *SignUpResponse {
	return &SignUpResponse{
		Success: true,
		Message: "회원가입이 정상처리되었습니다.",
		Token:   token,
	}
}

func FailSignUp(message string, err error) *SignUpResponse {
	return &SignUpResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
