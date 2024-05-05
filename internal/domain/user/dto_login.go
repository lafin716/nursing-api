package user

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,min=1,max=100"`
	Password string `json:"password" validate:"required,password"`
}
