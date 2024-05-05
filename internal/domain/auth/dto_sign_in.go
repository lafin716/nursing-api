package auth

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email,min=1,max=255" example:"lafin716@gmail.com"`
	Password string `json:"password" validate:"required,password" example:"lafin1234"`
}
