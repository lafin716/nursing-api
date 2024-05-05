package user

type CheckEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}
