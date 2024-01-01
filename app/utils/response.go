package utils

import "github.com/gofiber/fiber/v2"

type ResponseStatus struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}

func CreateResponseStatus(code int, params ...string) *ResponseStatus {
	return &ResponseStatus{
		Code:    code,
		Message: GetMessage(code, params...),
	}
}

func ResponseData(status *ResponseStatus, data interface{}) interface{} {
	resp := fiber.Map{
		"code":    status.Code,
		"message": status.Message,
	}

	if data != nil {
		resp["data"] = data
	}

	return resp
}

func ResponseError(status *ResponseStatus, errors any) interface{} {
	resp := fiber.Map{
		"code":    status.Code,
		"message": status.Message,
	}

	if errors != nil {
		resp["errors"] = errors
	}

	return resp
}
