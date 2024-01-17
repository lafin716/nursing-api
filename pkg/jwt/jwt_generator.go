package jwt

import (
	"crypto/sha256"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtGenerator struct {
	config *JwtConfig
}

type TokenRequest struct {
	ID          string
	Credentials []string
}

type Tokens struct {
	AccessToken         string `json:"access_token"`
	AccessTokenExpires  int64  `json:"access_token_expires"`
	RefreshToken        string `json:"refresh_token"`
	RefreshTokenExpires int64  `json:"refresh_token_expires"`
}

func NewJwtGenerator(config *JwtConfig) *JwtGenerator {
	return &JwtGenerator{
		config: config,
	}
}

func (g *JwtGenerator) GenerateNewTokens(request *TokenRequest) (*Tokens, error) {
	accessToken, accessExpires, err := g.generateNewAccessToken(request.ID, request.Credentials)
	if err != nil {
		return nil, err
	}

	refreshToken, refreshExpires, err := g.generateNewRefreshToken(request.ID)
	if err != nil {
		return nil, err
	}

	return &Tokens{
		AccessToken:         accessToken,
		AccessTokenExpires:  accessExpires,
		RefreshToken:        refreshToken,
		RefreshTokenExpires: refreshExpires,
	}, nil
}

func (g *JwtGenerator) generateNewAccessToken(id string, credentials []string) (string, int64, error) {
	secret := g.config.SecretKey
	minutesCount := g.config.AccessTokenExpires
	accessTokenExpires := time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()
	claims := jwt.MapClaims{
		"id":      id,
		"expires": accessTokenExpires,
	}

	// TODO 현재는 사용안함 추후 권한 관리 추가시 활성화
	//for _, credentials := range credentials {
	//	claims[credentials] = true
	//}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", 0, err
	}

	return t, accessTokenExpires, nil
}

func (g *JwtGenerator) generateNewRefreshToken(id string) (string, int64, error) {
	secret := g.config.RefreshKey
	hash := sha256.New()
	refresh := g.config.RefreshKey + time.Now().String()

	_, err := hash.Write([]byte(refresh))
	if err != nil {
		return "", 0, nil
	}

	hoursCount := g.config.RefreshTokenExpires
	refreshExpireTime := time.Now().Add(time.Hour * time.Duration(hoursCount)).Unix()
	claims := jwt.MapClaims{
		"id":      id,
		"expires": refreshExpireTime,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", 0, err
	}

	return t, refreshExpireTime, nil
}
