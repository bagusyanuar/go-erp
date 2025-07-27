package handler

import (
	"log"

	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/internal/delivery/request"
	"github.com/bagusyanuar/go-erp/internal/service"
	"github.com/bagusyanuar/go-erp/pkg/exception"
	"github.com/bagusyanuar/go-erp/pkg/lib"
	"github.com/bagusyanuar/go-erp/pkg/lib/response"
	"github.com/gofiber/fiber/v2"
)

type MaterialInventoryAdjustmentHandler struct {
	MaterialInventoryAdjustmentService service.MaterialInventoryAdjustmentService
	Config                             *config.AppConfig
}

func NewMaterialInventoryAdjustmentHandler(
	materialInventoryAdjustmentService service.MaterialInventoryAdjustmentService,
	cfg *config.AppConfig,
) *MaterialInventoryAdjustmentHandler {
	return &MaterialInventoryAdjustmentHandler{
		MaterialInventoryAdjustmentService: materialInventoryAdjustmentService,
		Config:                             cfg,
	}
}

func (c *MaterialInventoryAdjustmentHandler) Create(ctx *fiber.Ctx) error {
	request := new(request.MaterialInventoryAdjustmentSchema)
	if err := ctx.BodyParser(request); err != nil {
		log.Println(err.Error())
		return response.MakeResponseBadRequest(ctx, response.APIResponseOptions[any]{
			Message: exception.ErrBadRequest.Error(),
		})
	}

	messages, err := lib.Validate(c.Config.Validator, request)
	if err != nil {
		return response.MakeAPIResponseErrorValidation(ctx, messages)
	}

	res := c.MaterialInventoryAdjustmentService.Create(ctx.UserContext(), request)
	return response.MakeAPIResponseFromService(ctx, res)
}

func (c *MaterialInventoryAdjustmentHandler) FindAll(ctx *fiber.Ctx) error {
	queryParams := new(request.MaterialInventoryAdjustmetQuery)
	if err := ctx.QueryParser(queryParams); err != nil {
		return response.MakeResponseBadRequest(ctx, response.APIResponseOptions[any]{
			Message: exception.ErrBadRequest.Error(),
		})
	}

	messages, err := lib.Validate(c.Config.Validator, queryParams)
	if err != nil {
		return response.MakeAPIResponseErrorValidation(ctx, messages)
	}

	res := c.MaterialInventoryAdjustmentService.FindAll(ctx.UserContext(), queryParams)
	return response.MakeAPIResponseFromService(ctx, res)
}
