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

	return c.Status(fiber.StatusInternalServerError).JSON(utils.ResponseError(
		utils.CreateResponseStatus(code),
		err.Error(),
	))
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(utils.ResponseError(
		utils.CreateResponseStatus(utils.CODE_NOTFOUND),
		nil,
	))
}
