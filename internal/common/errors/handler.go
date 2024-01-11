package errors

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/common/response"
)

func GlobalErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return response.New(code).
		SetMessage(err.Error()).
		SetErrors(err).
		Error(c)
}
