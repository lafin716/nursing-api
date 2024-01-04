package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"strings"
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
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewJwtGenerator(config *JwtConfig) *JwtGenerator {
	return &JwtGenerator{
		config: config,
	}
}

func (g *JwtGenerator) GenerateNewTokens(request *TokenRequest) (*Tokens, error) {
	accessToken, err := g.generateNewAccessToken(request.ID, request.Credentials)
	if err != nil {
		return nil, err
	}

	refreshToken, err := g.generateNewRefreshToken()
	if err != nil {
		return nil, err
	}

	return &Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (g *JwtGenerator) generateNewAccessToken(id string, credentials []string) (string, error) {
	secret := g.config.SecretKey
	minutesCount := g.config.AccessTokenExpires
	claims := jwt.MapClaims{
		"id":      id,
		"expires": time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix(),
	}

	// TODO 현재는 사용안함 추후 권한 관리 추가시 활성화
	//for _, credentials := range credentials {
	//	claims[credentials] = true
	//}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (g *JwtGenerator) generateNewRefreshToken() (string, error) {
	hash := sha256.New()
	refresh := g.config.RefreshKey + time.Now().String()

	_, err := hash.Write([]byte(refresh))
	if err != nil {
		return "", nil
	}

	hoursCount := g.config.RefreshTokenExpires
	expireTime := fmt.Sprint(time.Now().Add(time.Hour * time.Duration(hoursCount)).Unix())
	t := hex.EncodeToString(hash.Sum(nil)) + "." + expireTime

	return t, nil
}

// TODO 이 함수는 필요는 해보이나 일단 당장은 사용하지 않으므로 대기
func ParseRefreshToken(refreshToken string) (int64, error) {
	return strconv.ParseInt(strings.Split(refreshToken, ".")[1], 0, 64)
}
