package handler

import (
	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/internal/delivery/request"
	"github.com/bagusyanuar/go-erp/internal/service"
	"github.com/bagusyanuar/go-erp/pkg/exception"
	"github.com/bagusyanuar/go-erp/pkg/lib"
	"github.com/bagusyanuar/go-erp/pkg/lib/response"
	"github.com/gofiber/fiber/v2"
)

type MaterialInventoryHandler struct {
	MaterialInventoryService service.MaterialInventoryService
	Config                   *config.AppConfig
}

func NewMaterialInventoryHandler(materialInventoryService service.MaterialInventoryService, cfg *config.AppConfig) *MaterialInventoryHandler {
	return &MaterialInventoryHandler{
		MaterialInventoryService: materialInventoryService,
		Config:                   cfg,
	}
}

func (c *MaterialInventoryHandler) Create(ctx *fiber.Ctx) error {
	// token := ctx.Locals("user").(*jwt.Token) // casting ke *jwt.Token
	// claims := token.Claims.(jwt.MapClaims)   // baru ambil claim-nya
	// userID := claims["sub"].(string)

	request := new(request.MaterialInventorySchema)
	if err := ctx.BodyParser(request); err != nil {
		return response.MakeResponseBadRequest(ctx, response.APIResponseOptions[any]{
			Message: exception.ErrBadRequest.Error(),
		})
	}

	messages, err := lib.Validate(c.Config.Validator, request)
	if err != nil {
		return response.MakeAPIResponseErrorValidation(ctx, messages)
	}

	res := c.MaterialInventoryService.Create(ctx.UserContext(), request)
	return response.MakeAPIResponseFromService(ctx, res)
}
