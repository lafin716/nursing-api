package middleware

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/common/response"
	"nursing_api/internal/domain/auth"
	"nursing_api/pkg/jwt"
	"time"
)

type TokenVerifyMiddleware struct {
	tokenRepo auth.Repository
	jwtClient *jwt.JwtClient
}

func NewJwtMiddleware(
	tokenRepo auth.Repository,
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
			return response.New(response.CODE_INVALID_JWT).SetErrors(err).Error(c)
		}

		userId := claims.UserID
		token, err := t.tokenRepo.GetToken(userId)
		if err != nil {
			return response.New(response.CODE_FAIL_SIGN_UP).SetErrors(err).Error(c)
		}

		// 토큰 만료 여부
		expire := time.Unix(token.AccessTokenExpires, 0)
		isExpireAccessToken := expire.Before(time.Now())

		// 토큰이 만료된 경우
		if isExpireAccessToken {
			return response.New(response.CODE_EXPIRED_ACCESSTOKEN).Error(c)
		}

		return c.Next()
	}
}
