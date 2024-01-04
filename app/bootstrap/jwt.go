package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/app/utils"
	"nursing_api/pkg/auth"
	"os"
	"strconv"
)

func GetJwtClient() *auth.JwtClient {
	accessTokenExpires, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRES"))
	refreshTokenExpires, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRES"))

	jwtConfig := &auth.JwtConfig{
		SecretKey:           os.Getenv("JWT_SECRET_KEY"),
		RefreshKey:          os.Getenv("REFRESH_SECRET_KEY"),
		AccessTokenExpires:  accessTokenExpires,
		RefreshTokenExpires: refreshTokenExpires,
		ErrorHandler:        jwtErrorHandler,
	}

	return auth.NewJwtClient(jwtConfig)
}

func jwtErrorHandler(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return utils.ResponseEntity{
			Code:   utils.CODE_INVALID_JWT,
			Errors: err.Error(),
		}.ResponseBadRequest(c)
	}

	return utils.ResponseEntity{
		Code:   utils.CODE_UNAUTHORIZED,
		Errors: err.Error(),
	}.ResponseUnAuthorized(c)
}
