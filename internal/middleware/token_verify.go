package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/domain/auth"
	"nursing_api/pkg/jwt"
	"time"
)

type TokenVerifyMiddleware struct {
	tokenRepo auth.AuthRepository
	jwtClient *jwt.JwtClient
}

func NewJwtMiddleware(
	tokenRepo auth.AuthRepository,
	jwtClient *jwt.JwtClient,
) *TokenVerifyMiddleware {
	return &TokenVerifyMiddleware{
		jwtClient: jwtClient,
		tokenRepo: tokenRepo,
	}
}

func (t *TokenVerifyMiddleware) Resolve() []fiber.Handler {
	return []func(*fiber.Ctx) error{
		t.jwtClient.Middleware.JwtProtected(),
		t.expireProtected(),
	}
}

func (t *TokenVerifyMiddleware) expireProtected() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		claims, err := t.jwtClient.Parser.ExtractTokenMetadata(c)
		if err != nil {
			return INVALID_TOKEN
		}

		userId := claims.UserID
		token, err := t.tokenRepo.GetToken(userId)
		if err != nil {
			return NO_MATCHING
		}

		// 토큰 만료 여부
		expire := time.Unix(token.AccessTokenExpires, 0)
		isExpireAccessToken := expire.Before(time.Now())

		// 자동로그인을 사용하지 않는 경우 오류 응답
		if !token.AutoLogin && isExpireAccessToken {
			return EXPIRE_TOKEN
		}

		// 자동 로그인인 경우 Refresh Token 검증 후 토큰 갱신처리
		if token.AutoLogin && isExpireAccessToken {
			refreshToken, err := t.jwtClient.Parser.ExtractRefreshTokenMetadata(c)
			if err != nil {
				return INVALID_TOKEN
			}
			if userId != refreshToken.UserID {
				return INVALID_TOKEN
			}

			refreshExpire := time.Unix(refreshToken.Expires, 0)
			isExpireRefreshToken := refreshExpire.Before(time.Now())
			if isExpireRefreshToken {
				return EXPIRE_TOKEN
			}

			newToken, err := t.jwtClient.Generator.GenerateNewTokens(&jwt.TokenRequest{
				ID:          refreshToken.UserID.String(),
				Credentials: []string{},
			})
			if err != nil {
				return FAIL_CREATE_TOKEN
			}

			_, err = t.tokenRepo.DeleteToken(userId)
			if err != nil {
				return FAIL_CREATE_TOKEN
			}
			_, err = t.tokenRepo.SaveToken(userId, &auth.Token{
				AccessToken:         newToken.AccessToken,
				RefreshToken:        newToken.RefreshToken,
				AccessTokenExpires:  newToken.AccessTokenExpires,
				RefreshTokenExpires: newToken.RefreshTokenExpires,
				AutoLogin:           token.AutoLogin,
			})
			if err != nil {
				return FAIL_CREATE_TOKEN
			}
		}

		return c.Next()
	}
}

var (
	INVALID_TOKEN     = errors.New("유효하지 않은 토큰입니다.")
	EXPIRE_TOKEN      = errors.New("토큰이 만료되었습니다. 다시 로그인 해주세요.")
	FAIL_CREATE_TOKEN = errors.New("토큰 발급에 실패하였습니다.")
	NO_MATCHING       = errors.New("로그인 정보가 없습니다. 다시 로그인 해주세요.")
)
