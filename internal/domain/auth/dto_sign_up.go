package auth

type SignUpRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=30"`
	Email    string `json:"email" validate:"required,email,min=1,max=100"`
	Password string `json:"password" validate:"required,password"`
}

type SignUpResponse struct {
	Success bool
	Message string
	Token   *Token
	Error   error
}

func OkSignUp(token *Token) *SignUpResponse {
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
