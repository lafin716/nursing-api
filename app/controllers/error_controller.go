package controllers

import (
	"nursing_api/app/container"
	"nursing_api/app/utils"
)

type FiberError struct {
	Code    int
	Message string
}

func NotFound(container *container.FiberContainer) error {
	c := container.Ctx
	return utils.ResponseEntity{
		Code: utils.CODE_NOTFOUND,
	}.ResponseOk(c)
}
