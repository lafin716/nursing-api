package user

type CheckEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type CheckEmailResponse struct {
	Success bool
	Message string
	Error   error
}

func OkCheckEmail() *CheckEmailResponse {
	return &CheckEmailResponse{
		Success: true,
		Message: "사용 가능한 이메일입니다.",
	}
}

func FailCheckEmail(message string, err error) *CheckEmailResponse {
	return &CheckEmailResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
