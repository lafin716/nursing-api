package api

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/common/response"
)

func Ok(data interface{}, c *fiber.Ctx) error {
	return response.New(response.CODE_SUCCESS).
		SetData(data).
		Ok(c)
}

func Fail(message string, err interface{}, c *fiber.Ctx) error {
	return response.New(response.CODE_ERROR).
		SetMessage(message).
		SetErrors(err).
		Error(c)
}

func FailAuth(err interface{}, c *fiber.Ctx) error {
	return response.New(response.CODE_INVALID_JWT).
		SetErrors(err).
		Error(c)
}

func FailParam(err interface{}, c *fiber.Ctx) error {
	return response.New(response.CODE_INVALID_PARAM).
		SetErrors(err).
		Error(c)
}
