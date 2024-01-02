package utils

import "github.com/gofiber/fiber/v2"

type ResponseEntity struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:",data"`
	Errors  interface{} `json:",errors"`
}

func (entity *ResponseEntity) SetMessage() *ResponseEntity {
	entity.Message = GetMessage(entity.Code)
	return entity
}

func (entity *ResponseEntity) SetMessageWithParams(params ...string) *ResponseEntity {
	entity.Message = GetMessage(entity.Code, params...)
	return entity
}

func ResponseOk(c *fiber.Ctx, entity *ResponseEntity) error {
	return c.Status(fiber.StatusOK).JSON(entity)
}

func ResponseError(c *fiber.Ctx, entity *ResponseEntity) error {
	return c.Status(fiber.StatusInternalServerError).JSON(entity)
}

func ResponseBadRequest(c *fiber.Ctx, entity *ResponseEntity) error {
	return c.Status(fiber.StatusBadRequest).JSON(entity)
}

func ResponseUnAuthorized(c *fiber.Ctx, entity *ResponseEntity) error {
	return c.Status(fiber.StatusUnauthorized).JSON(entity)
}
