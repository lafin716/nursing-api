package api

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/common/dto"
	"nursing_api/internal/common/response"
)

func ResolveResponse[T any](result dto.BaseResponse[T], c *fiber.Ctx) error {
	if result.IsSuccess() {
		return response.New(result.GetResultCode()).
			SetData(result.GetData()).
			Ok(c)
	}
	return response.New(result.GetResultCode()).
		SetErrors(result.GetError()).
		Error(c)
}

func Ok(data interface{}, c *fiber.Ctx) error {
	return response.New(response.CODE_SUCCESS).
		SetData(data).
		Ok(c)
}

func OkWithMessage(message string, data interface{}, c *fiber.Ctx) error {
	return response.New(response.CODE_SUCCESS).
		SetMessage(message).
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
