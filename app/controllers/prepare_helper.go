package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"nursing_api/app/utils"
	"nursing_api/pkg/auth"
	"nursing_api/pkg/validates"
)

// 파라미터 파싱
func parseParams(c *fiber.Ctx, req interface{}) error {
	if err := c.BodyParser(req); err != nil {
		code := fiber.StatusInternalServerError

		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		_ = utils.ResponseEntity{
			Code:   code,
			Errors: e,
		}.ResponseBadRequest(c)
		return err
	}

	return nil
}

// 파라미터 객체 유효성 검사
func validate(c *fiber.Ctx, model interface{}) error {
	validate := validates.NewValidator()
	if errs := validate.Struct(model); errs != nil {
		err := validates.ValidateErrors(errs)
		_ = utils.ResponseEntity{
			Code:   utils.CODE_INVALID_PARAM,
			Errors: err,
		}.ResponseBadRequest(c)
		return errs
	}

	return nil
}

// jwt 토큰 파싱
func parseJwt(c *fiber.Ctx) (*auth.TokenMetadata, error) {
	claims, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		_ = utils.ResponseEntity{
			Code:   utils.CODE_INVALID_JWT,
			Errors: err,
		}.ResponseBadRequest(c)
		return nil, err
	}

	return claims, nil
}
