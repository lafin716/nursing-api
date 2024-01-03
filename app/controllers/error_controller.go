package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"nursing_api/app/utils"
)

type FiberError struct {
	Code    int
	Message string
}

func GlobalErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return utils.ResponseEntity{
		Code:   code,
		Errors: err.Error(),
	}.ResponseError(c)
}

func NotFound(c *fiber.Ctx) error {
	return utils.ResponseEntity{
		Code: utils.CODE_NOTFOUND,
	}.ResponseOk(c)
}
