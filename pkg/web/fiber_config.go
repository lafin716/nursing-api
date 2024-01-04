package web

import "github.com/gofiber/fiber/v2"

type FiberConfig struct {
	AppName      string
	ErrorHandler func(c *fiber.Ctx, err error) error
}

func NewFiberConfig(appName string) *FiberConfig {
	return &FiberConfig{
		AppName: appName,
	}
}

func (c *FiberConfig) SetErrorHandler(handler func(c *fiber.Ctx, err error) error) {
	c.ErrorHandler = handler
}
