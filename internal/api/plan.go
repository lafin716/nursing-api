package api

import (
	"github.com/gofiber/fiber/v2"
	"nursing_api/internal/domain/plan"
)

type PlanHttpApi interface {
	Today(ctx *fiber.Ctx) error
	Summary(ctx *fiber.Ctx) error
	TakePlan(ctx *fiber.Ctx) error
	TakePill(ctx *fiber.Ctx) error
	UpdateMemo(ctx *fiber.Ctx) error
}

type planHttpApi struct {
	service plan.PlanUseCase
}

func (p planHttpApi) Today(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (p planHttpApi) Summary(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (p planHttpApi) TakePlan(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (p planHttpApi) TakePill(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (p planHttpApi) UpdateMemo(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func NewPlanHttpApi(
	service plan.PlanUseCase,
) PlanHttpApi {
	return &planHttpApi{
		service: service,
	}
}
