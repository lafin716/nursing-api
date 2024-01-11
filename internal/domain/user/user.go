package user

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID `json:"id" validate:"required,uuid"`
	Name         string    `json:"name" validate:"required,lte=100"`
	Email        string    `json:"email" validate:"required,email,lte=255"`
	PasswordHash string    `json:"-" validate:"required,lte=255"`
	UserStatus   string    `json:"user_status" validate:"required"`
	UserType     string    `json:"user_type"`
	FailCount    int       `json:"fail_count"`
	LastLoggedIn time.Time `json:"last_logged_in"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserRepository interface {
	CreateUser(user *User) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserById(userId uuid.UUID) (*User, error)
	CountUserByEmail(email string) (int, error)
}

type UserUseCase interface {
	RegisterUser(req *RegisterRequest) *RegisterResponse
	VerifyUser(req *LoginRequest) *LoginResponse
	GetUser(userId uuid.UUID) *GetUserResponse
}
