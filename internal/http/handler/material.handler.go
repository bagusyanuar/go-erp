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

type MaterialHandler struct {
	MaterialService service.MaterialService
	Config          *config.AppConfig
}

func NewMaterialHandler(materialService service.MaterialService, cfg *config.AppConfig) *MaterialHandler {
	return &MaterialHandler{
		MaterialService: materialService,
		Config:          cfg,
	}
}

func (c *MaterialHandler) FindAll(ctx *fiber.Ctx) error {
	queryParams := new(request.MaterialQuery)
	if err := ctx.QueryParser(queryParams); err != nil {
		return response.MakeResponseBadRequest(ctx, response.APIResponseOptions[any]{
			Message: exception.ErrBadRequest.Error(),
		})
	}

	messages, err := lib.Validate(c.Config.Validator, queryParams)
	if err != nil {
		return response.MakeAPIResponseErrorValidation(ctx, messages)
	}

	res := c.MaterialService.FindAll(ctx.UserContext(), queryParams)
	return response.MakeAPIResponseFromService(ctx, res)
}

func (c *MaterialHandler) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	res := c.MaterialService.FindByID(ctx.UserContext(), id)
	return response.MakeAPIResponseFromService(ctx, res)
}

func (c *MaterialHandler) Create(ctx *fiber.Ctx) error {
	request := new(request.MaterialSchema)
	if err := ctx.BodyParser(request); err != nil {
		return response.MakeResponseBadRequest(ctx, response.APIResponseOptions[any]{
			Message: exception.ErrBadRequest.Error(),
		})
	}

	messages, err := lib.Validate(c.Config.Validator, request)
	if err != nil {
		return response.MakeAPIResponseErrorValidation(ctx, messages)
	}

	res := c.MaterialService.Create(ctx.UserContext(), request)
	return response.MakeAPIResponseFromService(ctx, res)
}
