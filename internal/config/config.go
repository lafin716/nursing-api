package config

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/common/errors"
	"nursing_api/internal/common/response"
	"nursing_api/pkg/database"
	"nursing_api/pkg/jwt"
	"nursing_api/pkg/web"
	"os"
	"strconv"
)

func NewFiberConfig() *web.FiberConfig {
	appName := os.Getenv("APP_NAME")
	config := web.NewFiberConfig(appName)
	config.SetErrorHandler(errors.GlobalErrorHandler)

	return config
}

func NewJwtConfig() *jwt.JwtConfig {
	accessTokenExpires, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRES"))
	if err != nil {
		return nil
	}
	refreshTokenExpires, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRES"))
	if err != nil {
		return nil
	}

	config := &jwt.JwtConfig{
		SecretKey:           os.Getenv("JWT_SECRET_KEY"),
		RefreshKey:          os.Getenv("REFRESH_SECRET_KEY"),
		AccessTokenExpires:  accessTokenExpires,
		RefreshTokenExpires: refreshTokenExpires,
		ErrorHandler:        defaultJwtErrorHandler,
	}

	return config
}

func defaultJwtErrorHandler(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return response.New(response.CODE_INVALID_JWT).
			SetErrors(err).
			BadRequest(c)
	}

	return response.New(response.CODE_UNAUTHORIZED).
		SetErrors(err).
		UnAuthorized(c)
}

func NewDatabaseConfig() *database.Config {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil
	}

	sslMode, err := strconv.ParseBool(os.Getenv("DB_SSL"))
	if err != nil {
		sslMode = false
	}

	dbDebug, err := strconv.ParseBool(os.Getenv("DB_DEBUG"))
	if err != nil {
		dbDebug = false
	}

	config := &database.Config{
		Host:      os.Getenv("DB_HOST"),
		Port:      port,
		Database:  os.Getenv("DB_NAME"),
		User:      os.Getenv("DB_USER"),
		Password:  os.Getenv("DB_PASSWORD"),
		SSLEnable: sslMode,
		Debug:     dbDebug,
	}

	return config
}
