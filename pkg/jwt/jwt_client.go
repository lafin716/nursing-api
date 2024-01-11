package jwt

import (
	"github.com/gofiber/fiber/v2"
)

type JwtConfig struct {
	SecretKey           string
	RefreshKey          string
	AccessTokenExpires  int
	RefreshTokenExpires int
	ErrorHandler        func(c *fiber.Ctx, err error) error
}

type JwtClient struct {
	Generator  *JwtGenerator
	Parser     *JwtParser
	Middleware *JwtMiddleware
}

func NewJwtClient(config *JwtConfig) *JwtClient {
	return &JwtClient{
		Generator:  NewJwtGenerator(config),
		Parser:     NewJwtParser(config),
		Middleware: NewJwtMiddleware(config),
	}
}
