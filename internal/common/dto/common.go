package dto

import "nursing_api/internal/common/response"

type baseResponse[T any] struct {
	Success    bool
	ResultCode response.ResultCode
	Code       int
	Message    string
	Data       *T
	Error      *error
}

type BaseResponse[T any] interface {
	IsSuccess() bool
	GetCode() int
	GetResultCode() response.ResultCode
	GetMessage() string
	GetData() *T
	GetError() *error
}

func (b *baseResponse[T]) IsSuccess() bool {
	return b.Success
}

func (b *baseResponse[T]) GetResultCode() response.ResultCode {
	return b.ResultCode
}

func (b *baseResponse[T]) GetCode() int {
	return b.Code
}

func (b *baseResponse[T]) GetMessage() string {
	return b.Message
}

func (b *baseResponse[T]) GetData() *T {
	return b.Data
}

func (b *baseResponse[T]) GetError() *error {
	return b.Error
}

func Ok[T any](code response.ResultCode, data *T) BaseResponse[T] {
	return &baseResponse[T]{
		Success:    true,
		ResultCode: code,
		Code:       int(code),
		Message:    response.GetMessage(code),
		Data:       data,
	}
}

func Fail[T any](code response.ResultCode, err error) BaseResponse[T] {
	return &baseResponse[T]{
		Success:    false,
		ResultCode: code,
		Code:       int(code),
		Message:    response.GetMessage(code),
		Error:      &err,
	}
}
