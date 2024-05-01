package errors

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/common/response"
)

type ErrorHttpApi interface {
	Handle(ctx *fiber.Ctx) error
}
type HTTP_ERROR_CODE int

func GlobalErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	msg := codeMessages[HTTP_ERROR_CODE(code)]
	if msg == "" {
		msg = codeMessages[INTERNAL_SERVER_ERROR]
	}

	return response.New(code).
		SetMessage(msg).
		SetErrors(err).
		CustomResponse(c)
}

const (
	NOTFOUND              = HTTP_ERROR_CODE(404)
	BAD_REQUEST           = HTTP_ERROR_CODE(400)
	UNAUTHRIZE            = HTTP_ERROR_CODE(401)
	INTERNAL_SERVER_ERROR = HTTP_ERROR_CODE(500)
)

var codeMessages = map[HTTP_ERROR_CODE]string{
	NOTFOUND:              "존재하지 않는 엔드포인트입니다.",
	BAD_REQUEST:           "올바른 요청이 아닙니다.",
	UNAUTHRIZE:            "권한이 없습니다.",
	INTERNAL_SERVER_ERROR: "오류가 발생하였습니다",
}
