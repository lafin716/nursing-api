package user

type GetUserResponse struct {
	Success bool
	Message string
	User    *User
	Error   error
}
