package domain

type BaseResponse[T interface{}] struct {
	Success bool
	Message string
	Data    T
	Error   error
}

func OkWithMsg[T interface{}](message string, data T) *BaseResponse[T] {
	return &BaseResponse[T]{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func Ok[T interface{}](data T) *BaseResponse[T] {
	return &BaseResponse[T]{
		Success: true,
		Data:    data,
	}
}

func Fail(message string, err error) *BaseResponse[any] {
	return &BaseResponse[any]{
		Success: false,
		Message: message,
		Error:   err,
	}
}
