package utils

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

	entity.Message = GetMessage(entity.Code)
}

func (entity ResponseEntity) ResponseOk(c *fiber.Ctx) error {
	resolveMessage(&entity)
	return c.Status(fiber.StatusOK).JSON(entity)
}

func (entity ResponseEntity) ResponseError(c *fiber.Ctx) error {
	resolveMessage(&entity)
	return c.Status(fiber.StatusInternalServerError).JSON(entity)
}

func (entity ResponseEntity) ResponseBadRequest(c *fiber.Ctx) error {
	resolveMessage(&entity)
	return c.Status(fiber.StatusBadRequest).JSON(entity)
}

func (entity ResponseEntity) ResponseUnAuthorized(c *fiber.Ctx) error {
	resolveMessage(&entity)
	return c.Status(fiber.StatusUnauthorized).JSON(entity)
}
