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

func (c *MaterialInventoryHandler) FindAll(ctx *fiber.Ctx) error {
	queryParams := new(request.MaterialInventoryQuery)
	if err := ctx.QueryParser(queryParams); err != nil {
		return response.MakeResponseBadRequest(ctx, response.APIResponseOptions[any]{
			Message: exception.ErrBadRequest.Error(),
		})
	}

	messages, err := lib.Validate(c.Config.Validator, queryParams)
	if err != nil {
		return response.MakeAPIResponseErrorValidation(ctx, messages)
	}

	res := c.MaterialInventoryService.FindAll(ctx.UserContext(), queryParams)
	return response.MakeAPIResponseFromService(ctx, res)
}

func (c *MaterialInventoryHandler) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	res := c.MaterialInventoryService.FindByID(ctx.UserContext(), id)
	return response.MakeAPIResponseFromService(ctx, res)
}

func (c *MaterialInventoryHandler) Create(ctx *fiber.Ctx) error {
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
