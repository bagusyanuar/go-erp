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

type MaterialCategoryHandler struct {
	MaterialCategoryService service.MaterialCategoryService
	Config                  *config.AppConfig
}

func NewMaterialCategoryHandler(materialCategoryService service.MaterialCategoryService, cfg *config.AppConfig) *MaterialCategoryHandler {
	return &MaterialCategoryHandler{
		MaterialCategoryService: materialCategoryService,
		Config:                  cfg,
	}
}

func (c *MaterialCategoryHandler) FindAll(ctx *fiber.Ctx) error {
	queryParams := new(request.MaterialCategoryQuery)
	if err := ctx.QueryParser(queryParams); err != nil {
		return response.MakeResponseBadRequest(ctx, response.APIResponseOptions[any]{
			Message: exception.ErrBadRequest.Error(),
		})
	}

	messages, err := lib.Validate(c.Config.Validator, queryParams)
	if err != nil {
		return response.MakeAPIResponseErrorValidation(ctx, messages)
	}

	res := c.MaterialCategoryService.FindAll(ctx.UserContext(), queryParams)
	return response.MakeAPIResponseFromService(ctx, res)
}

func (c *MaterialCategoryHandler) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	res := c.MaterialCategoryService.FindByID(ctx.UserContext(), id)
	return response.MakeAPIResponseFromService(ctx, res)
}

func (c *MaterialCategoryHandler) Create(ctx *fiber.Ctx) error {
	request := new(request.MaterialCategorySchema)
	if err := ctx.BodyParser(request); err != nil {
		return response.MakeResponseBadRequest(ctx, response.APIResponseOptions[any]{
			Message: exception.ErrBadRequest.Error(),
		})
	}

	messages, err := lib.Validate(c.Config.Validator, request)
	if err != nil {
		return response.MakeAPIResponseErrorValidation(ctx, messages)
	}

	res := c.MaterialCategoryService.Create(ctx.UserContext(), request)
	return response.MakeAPIResponseFromService(ctx, res)
}
