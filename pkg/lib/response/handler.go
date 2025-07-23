package response

import "github.com/gofiber/fiber/v2"

type (
	APIResponse[T any] struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    T      `json:"data,omitempty"`
		Meta    any    `json:"meta,omitempty"`
	}

	APIResponseOptions[T any] struct {
		Message string
		Data    T
		Meta    any
	}
)

func MakeAPIResponse[T any](ctx *fiber.Ctx, res APIResponse[T]) error {
	return ctx.Status(res.Code).JSON(res)
}

func MakeAPIResponseFromService[T any](ctx *fiber.Ctx, svcResponse ServiceResponse[T]) error {
	res := createResponseFromService(svcResponse)
	return ctx.Status(res.Code).JSON(res)
}

func createResponseFromService[T any](res ServiceResponse[T]) APIResponse[T] {
	return APIResponse[T]{
		Code:    res.Code,
		Message: res.Message,
		Data:    res.Data,
		Meta:    res.Meta,
	}
}

//code:2**
func MakeResponseOK[T any](ctx *fiber.Ctx, opts APIResponseOptions[T]) error {
	return ctx.Status(fiber.StatusOK).JSON(APIResponse[T]{
		Code:    fiber.StatusOK,
		Message: opts.Message,
		Data:    opts.Data,
		Meta:    opts.Meta,
	})
}

func MakeResponseCreated[T any](ctx *fiber.Ctx, opts APIResponseOptions[T]) error {
	return ctx.Status(fiber.StatusCreated).JSON(APIResponse[T]{
		Code:    fiber.StatusCreated,
		Message: opts.Message,
		Data:    opts.Data,
		Meta:    opts.Meta,
	})
}

//code:4**

func MakeResponseBadRequest[T any](ctx *fiber.Ctx, opts APIResponseOptions[T]) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(APIResponse[T]{
		Code:    fiber.StatusBadRequest,
		Message: opts.Message,
		Data:    opts.Data,
		Meta:    opts.Meta,
	})
}

func MakeResponseUnauthorized[T any](ctx *fiber.Ctx, opts APIResponseOptions[T]) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(APIResponse[T]{
		Code:    fiber.StatusUnauthorized,
		Message: opts.Message,
		Data:    opts.Data,
		Meta:    opts.Meta,
	})
}

func MakeResponseForbidden[T any](ctx *fiber.Ctx, opts APIResponseOptions[T]) error {
	return ctx.Status(fiber.StatusForbidden).JSON(APIResponse[T]{
		Code:    fiber.StatusForbidden,
		Message: opts.Message,
		Data:    opts.Data,
		Meta:    opts.Meta,
	})
}

func MakeResponseNotFound[T any](ctx *fiber.Ctx, opts APIResponseOptions[T]) error {
	return ctx.Status(fiber.StatusNotFound).JSON(APIResponse[T]{
		Code:    fiber.StatusNotFound,
		Message: opts.Message,
		Data:    opts.Data,
		Meta:    opts.Meta,
	})
}

func MakeResponseUnprocessableEntity[T any](ctx *fiber.Ctx, opts APIResponseOptions[T]) error {
	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(APIResponse[T]{
		Code:    fiber.StatusUnprocessableEntity,
		Message: opts.Message,
		Data:    opts.Data,
		Meta:    opts.Meta,
	})
}

//code:500
func MakeResponseInternalServerError[T any](ctx *fiber.Ctx, opts APIResponseOptions[T]) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(APIResponse[T]{
		Code:    fiber.StatusInternalServerError,
		Message: opts.Message,
		Data:    opts.Data,
		Meta:    opts.Meta,
	})
}
