package auth

import "nursing_api/pkg/jwt"

type RefreshTokenRequest struct {
	AccessToken  *jwt.TokenMetadata
	RefreshToken *jwt.TokenMetadata
}

type RefreshTokenResponse struct {
	Success bool
	Message string
	Token   *Token
	Error   error
}

func OkRefreshToken(token *Token) *RefreshTokenResponse {
	return &RefreshTokenResponse{
		Success: true,
		Message: "토큰이 갱신되었습니다.",
		Token:   token,
	}
}

func FailRefreshToken(message string, err error) *RefreshTokenResponse {
	return &RefreshTokenResponse{
		Success: false,
		Message: message,
		Error:   err,
	}
}
