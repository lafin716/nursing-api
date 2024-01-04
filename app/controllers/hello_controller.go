package controllers

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/app/container"
)

func Hello(c *container.FiberContainer) error {
	ctx := c.Ctx
	return ctx.Status(fiber.StatusAccepted).SendString("Hello world")
}
