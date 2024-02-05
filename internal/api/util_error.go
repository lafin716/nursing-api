package api

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/common/response"
)

func FailParam(err interface{}, c *fiber.Ctx) error {
	return response.New(response.CODE_INVALID_PARAM).
		SetErrors(err).
		Error(c)
}
