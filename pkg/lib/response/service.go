package response

import "github.com/gofiber/fiber/v2"

type (
	ServiceResponse[T any] struct {
		Code    int
		Error   error
		Message string
		Data    T
		Meta    any
	}

	ServiceResponseOptions[T any] struct {
		Error   error
		Message string
		Data    T
		Meta    any
	}
)

//code:2**
func ServiceOK[T any](opt ServiceResponseOptions[T]) ServiceResponse[T] {
	return ServiceResponse[T]{
		Code:    fiber.StatusOK,
		Error:   nil,
		Message: opt.Message,
		Data:    opt.Data,
		Meta:    opt.Meta,
	}
}

func ServiceCreated[T any](opt ServiceResponseOptions[T]) ServiceResponse[T] {
	return ServiceResponse[T]{
		Code:    fiber.StatusCreated,
		Error:   nil,
		Message: opt.Message,
		Data:    opt.Data,
		Meta:    opt.Meta,
	}
}

//code:4**
func ServiceBadRequest[T any](opt ServiceResponseOptions[T]) ServiceResponse[T] {
	return ServiceResponse[T]{
		Code:    fiber.StatusBadRequest,
		Error:   opt.Error,
		Message: opt.Message,
		Data:    opt.Data,
		Meta:    opt.Meta,
	}
}

func ServiceUnauthorized[T any](opt ServiceResponseOptions[T]) ServiceResponse[T] {
	return ServiceResponse[T]{
		Code:    fiber.StatusUnauthorized,
		Error:   opt.Error,
		Message: opt.Message,
		Data:    opt.Data,
		Meta:    opt.Meta,
	}
}

func ServiceForbidden[T any](opt ServiceResponseOptions[T]) ServiceResponse[T] {
	return ServiceResponse[T]{
		Code:    fiber.StatusForbidden,
		Error:   opt.Error,
		Message: opt.Message,
		Data:    opt.Data,
		Meta:    opt.Meta,
	}
}

func ServiceNotFound[T any](opt ServiceResponseOptions[T]) ServiceResponse[T] {
	return ServiceResponse[T]{
		Code:    fiber.StatusNotFound,
		Error:   opt.Error,
		Message: opt.Message,
		Data:    opt.Data,
		Meta:    opt.Meta,
	}
}

func ServiceUnprocessableEntity[T any](opt ServiceResponseOptions[T]) ServiceResponse[T] {
	return ServiceResponse[T]{
		Code:    fiber.StatusUnprocessableEntity,
		Error:   opt.Error,
		Message: opt.Message,
		Data:    opt.Data,
		Meta:    opt.Meta,
	}
}

//code:500
func ServiceInternalServerError[T any](opt ServiceResponseOptions[T]) ServiceResponse[T] {
	return ServiceResponse[T]{
		Code:    fiber.StatusInternalServerError,
		Error:   opt.Error,
		Message: opt.Message,
		Data:    opt.Data,
		Meta:    opt.Meta,
	}
}
