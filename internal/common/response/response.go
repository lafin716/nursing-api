package response

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

type ResponseEntity struct {
	Code          int         `json:"code"`
	Message       string      `json:"message,omitempty"`
	MessageParams []string    `json:"-"`
	Data          interface{} `json:"data,omitempty"`
	Errors        interface{} `json:"errors,omitempty"`
}

func resolveMessage(entity *ResponseEntity) {
	if strings.TrimSpace(entity.Message) != "" {
		return
	}

	if len(entity.MessageParams) > 0 {
		entity.Message = GetMessage(entity.Code, entity.MessageParams...)
	}

	if strings.TrimSpace(entity.Message) == "" {
		entity.Message = GetMessage(entity.Code)
	}
}

func New(code int) *ResponseEntity {
	return &ResponseEntity{
		Code: code,
	}
}

func (entity *ResponseEntity) SetMessage(message string) *ResponseEntity {
	entity.Message = message
	return entity
}

func (entity *ResponseEntity) SetMessageParams(params []string) *ResponseEntity {
	entity.MessageParams = params
	return entity
}

func (entity *ResponseEntity) SetData(data interface{}) *ResponseEntity {
	entity.Data = data
	return entity
}

func (entity *ResponseEntity) SetErrors(errs interface{}) *ResponseEntity {
	entity.Errors = errs
	return entity
}

func (entity *ResponseEntity) Ok(c *fiber.Ctx) error {
	resolveMessage(entity)
	return c.Status(fiber.StatusOK).JSON(entity)
}

func (entity *ResponseEntity) Error(c *fiber.Ctx) error {
	resolveMessage(entity)
	return c.Status(fiber.StatusInternalServerError).JSON(entity)
}

func (entity *ResponseEntity) BadRequest(c *fiber.Ctx) error {
	resolveMessage(entity)
	return c.Status(fiber.StatusBadRequest).JSON(entity)
}

func (entity *ResponseEntity) UnAuthorized(c *fiber.Ctx) error {
	resolveMessage(entity)
	return c.Status(fiber.StatusUnauthorized).JSON(entity)
}
