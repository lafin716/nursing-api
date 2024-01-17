package api

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/common/response"
)

type MainHttpApi interface {
	HelloWorld(ctx *fiber.Ctx) error
}

type mainHttpApi struct {
}

func NewMainHttpApi() MainHttpApi {
	return &mainHttpApi{}
}

func (m mainHttpApi) HelloWorld(ctx *fiber.Ctx) error {
	return response.New(response.CODE_SUCCESS).
		SetMessage("간병관리 API 서비스에 오신걸 환영합니다~!").
		Ok(ctx)
}
