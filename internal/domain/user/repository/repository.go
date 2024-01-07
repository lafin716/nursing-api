package repository

import (
	"github.com/google/uuid"
	"nursing_api/internal/domain/user"
)

type UserRepository interface {
	CreateUser(user *user.User) (*user.User, error)
	GetUserByEmail(email string) (*user.User, error)
	GetUserById(userId uuid.UUID) (*user.User, error)
	CountUserByEmail(email string) (int, error)
}
