package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strconv"
	"strings"
	"time"
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GenerateNewTokens(id string, credentials []string) (*Tokens, error) {
	accessToken, err := generateNewAccessToken(id, credentials)
	if err != nil {
		return nil, err
	}

	refreshToken, err := generateNewRefreshToken()
	if err != nil {
		return nil, err
	}

	return &Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func ParseRefreshToken(refreshToken string) (int64, error) {
	return strconv.ParseInt(strings.Split(refreshToken, ".")[1], 0, 64)
}

func generateNewAccessToken(id string, credentials []string) (string, error) {
	secret := os.Getenv("JWT_SECRET_KEY")
	minutesCount, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRES"))
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

func generateNewRefreshToken() (string, error) {
	hash := sha256.New()
	refresh := os.Getenv("REFRESH_SECRET_KEY") + time.Now().String()

	_, err := hash.Write([]byte(refresh))
	if err != nil {
		return "", nil
	}

	hoursCount, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRES"))
	expireTime := fmt.Sprint(time.Now().Add(time.Hour * time.Duration(hoursCount)).Unix())
	t := hex.EncodeToString(hash.Sum(nil)) + "." + expireTime

	return t, nil
}
