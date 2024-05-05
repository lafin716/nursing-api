package user

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=1,max=30"`
	Email    string `json:"email" validate:"required,email,min=1,max=100"`
	Password string `json:"password" validate:"required,password"`
}
