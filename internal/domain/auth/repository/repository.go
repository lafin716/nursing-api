package repository

import (
	"github.com/google/uuid"
	"nursing_api/internal/domain/auth"
)

type AuthRepository interface {
	SaveToken(userId uuid.UUID, token *auth.Token) (*auth.Token, error)
	GetToken(userId uuid.UUID) (*auth.Token, error)
	DeleteToken(userId uuid.UUID) (bool, error)
}
