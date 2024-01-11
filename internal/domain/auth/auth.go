package auth

import (
	"github.com/google/uuid"
)

type Token struct {
	AccessToken         string `json:"access_token"`
	RefreshToken        string `json:"refresh_token"`
	AccessTokenExpires  int64  `json:"access_token_expires"`
	RefreshTokenExpires int64  `json:"refresh_token_expires"`
}

type AuthRepository interface {
	SaveToken(userId uuid.UUID, token *Token) (*Token, error)
	GetToken(userId uuid.UUID) (*Token, error)
	DeleteToken(userId uuid.UUID) (bool, error)
}

type AuthUseCase interface {
	SignIn(req *SignInRequest) *SignInResponse
	SignUp(req *SignUpRequest) *SignUpResponse
}
