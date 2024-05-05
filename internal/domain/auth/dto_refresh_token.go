package auth

import "nursing_api/pkg/jwt"

type RefreshTokenRequest struct {
	AccessToken  *jwt.TokenMetadata
	RefreshToken *jwt.TokenMetadata
}
