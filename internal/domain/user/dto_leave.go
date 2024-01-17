package user

type LeaveResponse struct {
	Success bool
	Message string
	Error   error
}

func OkLeave() *LeaveResponse {
	return &LeaveResponse{
		Success: true,
		Message: "회원탈퇴가 정상적으로 처리되었습니다",
	}
}

func FailLeave(message string, err error) *LeaveResponse {
	return &LeaveResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
