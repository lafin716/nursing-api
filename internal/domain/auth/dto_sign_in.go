package auth

type SignInRequest struct {
	Email     string `json:"email" validate:"required,email,min=1,max=255" example:"lafin716@gmail.com"`
	Password  string `json:"password" validate:"required,password" example:"lafin1234"`
	AutoLogin bool   `json:"auto_login" example:"false"`
}

type SignInResponse struct {
	Success bool
	Message string
	Token   *Token
	Error   error
}

func OkSignIn(token *Token) *SignInResponse {
	return &SignInResponse{
		Success: true,
		Message: "토큰이 발급되었습니다.",
		Token:   token,
	}
}

func FailSignIn(message string, err error) *SignInResponse {
	return &SignInResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
