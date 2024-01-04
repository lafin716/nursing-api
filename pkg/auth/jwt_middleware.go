package auth

import (
	"github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

type JwtMiddleware struct {
	config *JwtConfig
}

func NewJwtMiddleware(config *JwtConfig) *JwtMiddleware {
	return &JwtMiddleware{
		config: config,
	}
}

func (m *JwtMiddleware) JwtProtected() func(*fiber.Ctx) error {
	config := jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(m.config.SecretKey)},
		ContextKey:   "jwt",
		ErrorHandler: m.config.ErrorHandler,
	}

	return jwtware.New(config)
}
