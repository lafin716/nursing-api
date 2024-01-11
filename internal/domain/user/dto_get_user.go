package user

type GetUserResponse struct {
	Success bool
	Message string
	User    *User
	Error   error
}

func OkGetUserResponse(user *User) *GetUserResponse {
	return &GetUserResponse{
		Success: true,
		Message: "유저 정보가 조회되었습니다.",
		User:    user,
	}
}

func FailGetUserResponse(message string, err error) *GetUserResponse {
	return &GetUserResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
