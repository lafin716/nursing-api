package api

import "github.com/gofiber/fiber/v2"

type ParamType int

const (
	QUERY = 0
	PARAM = 1
	BODY  = 2
)

type param struct {
	paramType ParamType
	ctx       *fiber.Ctx
}

func NewParam() *param {
	return new(param)
}

func resolveParam(paramType ParamType, ctx *fiber.Ctx) {

}
