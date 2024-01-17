package jwt

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
	token, err := p.verifyAccessToken(c)
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

func (p *JwtParser) ExtractRefreshTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := p.verifyRefreshToken(c)
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

func (p *JwtParser) verifyAccessToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := p.extractToken("Authorization", c)
	token, err := jwt.Parse(tokenString, p.jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (p *JwtParser) verifyRefreshToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := p.extractToken("RefreshToken", c)
	token, err := jwt.Parse(tokenString, p.jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (p *JwtParser) extractToken(headerName string, c *fiber.Ctx) string {
	bearerToken := c.Get(headerName)
	onlyToken := strings.Split(bearerToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func (p *JwtParser) jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(p.config.SecretKey), nil
}
