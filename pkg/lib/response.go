package lib

import (
	"github.com/bagusyanuar/go-erp/pkg/myexception"
	"github.com/gofiber/fiber/v2"
)

type StatusCode int

const (
	OK                  StatusCode = 200
	Created             StatusCode = 201
	BadRequest          StatusCode = 400
	Unauthorized        StatusCode = 401
	Forbidden           StatusCode = 403
	NotFound            StatusCode = 404
	UnprocessableEntity StatusCode = 422
	InternalServerError StatusCode = 500
)

type APIResponse[T any] struct {
	Code    StatusCode `json:"code"`
	Message string     `json:"message"`
	Data    T          `json:"data,omitempty"`
	Meta    any        `json:"meta,omitempty"`
}

type ResponseOptions[T any] struct {
	Message string
	Data    T
	Meta    any
}

func FromService[T any](service ServiceResponse[T]) APIResponse[T] {
	return APIResponse[T]{
		Code:    service.Status,
		Message: service.Message,
		Data:    service.Data,
		Meta:    service.Meta,
	}
}

func MakeResponse[T any](ctx *fiber.Ctx, response APIResponse[T]) error {
	return ctx.Status(int(response.Code)).JSON(response)
}

func MakeResponseFromService[T any](ctx *fiber.Ctx, service ServiceResponse[T]) error {
	response := FromService(service)
	return ctx.Status(int(response.Code)).JSON(response)
}
func ResponseErrValidation(ctx *fiber.Ctx, messages map[string][]string) error {
	return ResponseUnproccesableEntity(ctx, ResponseOptions[map[string][]string]{
		Message: myexception.ErrUnprocessableEntity.Error(),
		Data:    messages,
	})
}

func ResponseOK[T any](ctx *fiber.Ctx, opts ResponseOptions[T]) error {
	return ctx.Status(int(OK)).JSON(APIResponse[T]{
		Code:    OK,
		Message: opts.Message,
		Data:    opts.Data,
		Meta:    opts.Meta,
	})
}

func ResponseCreated[T any](ctx *fiber.Ctx, opts ResponseOptions[T]) error {
	return ctx.Status(int(Created)).JSON(APIResponse[T]{
		Code:    Created,
		Message: opts.Message,
		Data:    opts.Data,
		Meta:    opts.Meta,
	})
}

func ResponseBadRequest[T any](ctx *fiber.Ctx, opts ResponseOptions[T]) error {
	return ctx.Status(int(BadRequest)).JSON(APIResponse[T]{
		Code:    BadRequest,
		Message: opts.Message,
	})
}

func ResponseUnproccesableEntity[T any](ctx *fiber.Ctx, opts ResponseOptions[T]) error {
	return ctx.Status(int(UnprocessableEntity)).JSON(APIResponse[T]{
		Code:    UnprocessableEntity,
		Message: opts.Message,
		Data:    opts.Data,
	})
}

//repository:response

type RepositoryResponse[T any] struct {
	Message string
	Error   error
	Data    T
	Meta    any
}

func MakeRepositoryError[T any](err error) RepositoryResponse[T] {
	return RepositoryResponse[T]{
		Message: err.Error(),
		Error:   err,
	}
}

func MakeRepositorySuccess[T any](data T, meta any) RepositoryResponse[T] {
	return RepositoryResponse[T]{
		Message: "successfull",
		Error:   nil,
		Data:    data,
		Meta:    meta,
	}
}

//service:response

type ServiceResponse[T any] struct {
	Status  StatusCode
	Error   error
	Message string
	Data    T
	Meta    any
}

type ServiceResponseOptions[T any] struct {
	Error   error
	Message string
	Data    T
	Meta    any
}

func ServiceOK[T any](opt ServiceResponseOptions[T]) ServiceResponse[T] {
	return ServiceResponse[T]{
		Status:  OK,
		Error:   nil,
		Message: opt.Message,
		Data:    opt.Data,
		Meta:    opt.Meta,
	}
}

func ServiceCreated[T any](opt ServiceResponseOptions[T]) ServiceResponse[T] {
	return ServiceResponse[T]{
		Status:  Created,
		Error:   nil,
		Message: opt.Message,
		Data:    opt.Data,
		Meta:    opt.Meta,
	}
}

func ServiceNotFound[T any](opt ServiceResponseOptions[T]) ServiceResponse[T] {
	return ServiceResponse[T]{
		Status:  NotFound,
		Error:   opt.Error,
		Message: opt.Message,
		Data:    opt.Data,
		Meta:    opt.Meta,
	}
}

func ServiceInternalServerError[T any](opt ServiceResponseOptions[T]) ServiceResponse[T] {
	return ServiceResponse[T]{
		Status:  InternalServerError,
		Error:   opt.Error,
		Message: opt.Message,
		Data:    opt.Data,
		Meta:    opt.Meta,
	}
}
