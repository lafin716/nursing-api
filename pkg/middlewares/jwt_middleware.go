package middlewares

import (
	jwtMiddleware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"nursing_api/app/utils"
	"os"
)

func JwtProtected() func(*fiber.Ctx) error {
	config := jwtMiddleware.Config{
		SigningKey:   jwtMiddleware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET_KEY"))},
		ContextKey:   "jwt",
		ErrorHandler: jwtErrorHandler,
	}

	return jwtMiddleware.New(config)
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
