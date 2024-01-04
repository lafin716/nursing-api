package controllers

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/app/container"
)

type FiberHandler struct {
	Container *container.FiberContainer
}

func NewFiberHandler(container *container.FiberContainer) *FiberHandler {
	return &FiberHandler{
		Container: container,
	}
}

func (h *FiberHandler) Handle(controller func(container *container.FiberContainer) error) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		h.Container.Ctx = ctx
		err := controller(h.Container)
		return err
	}
}
