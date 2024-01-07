package dto

import "nursing_api/internal/domain/auth"

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}

type SignInResponse struct {
	Success bool
	Message string
	Token   *auth.Token
	Error   error
}

func OkSignIn(token *auth.Token) SignInResponse {
	return SignInResponse{
		Success: true,
		Message: "토큰이 발급되었습니다.",
		Token:   token,
	}
}

func FailSignIn(message string, err error) SignInResponse {
	return SignInResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
