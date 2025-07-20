package helper

import (
	"github.com/bagusyanuar/go-erp/internal/domain/dto"
	"github.com/gofiber/fiber/v2"
)

func MakeResponse[T any](ctx *fiber.Ctx, response dto.APIResponse[T]) error {
	return ctx.Status(int(response.Code)).JSON(response)
}
