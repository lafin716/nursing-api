package user

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=1,max=30"`
	Email    string `json:"email" validate:"required,email,min=1,max=100"`
	Password string `json:"password" validate:"required,password"`
}

type RegisterResponse struct {
	Success bool
	Message string
	User    *User
	Error   error
}

func OkRegisterUser(newUser *User) *RegisterResponse {
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
