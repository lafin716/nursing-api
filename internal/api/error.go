package api

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/common/response"
)

type ErrorHttpApi interface {
	Handle(ctx *fiber.Ctx) error
}

type HTTP_ERROR_CODE int
type errorHttpApi struct {
	code    HTTP_ERROR_CODE
	message string
}

func NewErrorHttpApi(code HTTP_ERROR_CODE) ErrorHttpApi {
	return &errorHttpApi{
		code:    code,
		message: codeMessages[code],
	}
}

func (e errorHttpApi) Handle(ctx *fiber.Ctx) error {
	return response.New(int(e.code)).
		SetMessage(e.message).
		Error(ctx)
}

const (
	NOTFOUND              = HTTP_ERROR_CODE(404)
	BAD_REQUEST           = HTTP_ERROR_CODE(400)
	UNAUTHRIZE            = HTTP_ERROR_CODE(401)
	INTERNAL_SERVER_ERROR = HTTP_ERROR_CODE(500)
)

var codeMessages = map[HTTP_ERROR_CODE]string{
	NOTFOUND:              "존재하지 않는 엔드포인트입니다.",
	UNAUTHRIZE:            "권한이 없습니다.",
	BAD_REQUEST:           "올바른 요청이 아닙니다.",
	INTERNAL_SERVER_ERROR: "오류가 발생하였습니다",
}
