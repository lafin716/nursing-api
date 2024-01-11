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
			return errors.New("유효하지 않은 토큰입니다.")
		}

		token, err := t.tokenRepo.GetToken(claims.UserID)
		if err != nil {
			return errors.New("로그인 정보가 없습니다. 다시 로그인 해주세요.")
		}

		expire := time.Unix(token.AccessTokenExpires, 0)
		if expire.Before(time.Now()) {
			return errors.New("토큰이 만료되었습니다. 다시 로그인 해주세요.")
		}

		return c.Next()
	}
}
