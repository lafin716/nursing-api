package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"strings"
)

type TokenMetadata struct {
	UserID     uuid.UUID
	Credential map[string]bool
	Expires    int64
}

type JwtParser struct {
	config *JwtConfig
}

func NewJwtParser(config *JwtConfig) *JwtParser {
	return &JwtParser{
		config: config,
	}
}

func (p *JwtParser) ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := p.verifyToken(c)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId, err := uuid.Parse(claims["id"].(string))
		if err != nil {
			return nil, err
		}

		expires := int64(claims["expires"].(float64))
		credentials := map[string]bool{}

		return &TokenMetadata{
			UserID:     userId,
			Credential: credentials,
			Expires:    expires,
		}, nil
	}

	return nil, err
}

func (p *JwtParser) verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := p.extractToken(c)
	token, err := jwt.Parse(tokenString, p.jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (p *JwtParser) extractToken(c *fiber.Ctx) string {
	bearerToken := c.Get("Authorization")
	onlyToken := strings.Split(bearerToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func (p *JwtParser) jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(p.config.SecretKey), nil
}
