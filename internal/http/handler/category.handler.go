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

type CategoryHandler struct {
	CategoryService service.CategoryService
	Config          *config.AppConfig
}

func NewCategoryHandler(categoryService service.CategoryService, cfg *config.AppConfig) *CategoryHandler {
	return &CategoryHandler{
		CategoryService: categoryService,
		Config:          cfg,
	}
}

func (c *CategoryHandler) FindAll(ctx *fiber.Ctx) error {
	queryParams := new(request.CategoryQuery)
	if err := ctx.QueryParser(queryParams); err != nil {
		return response.MakeResponseBadRequest(ctx, response.APIResponseOptions[any]{
			Message: exception.ErrBadRequest.Error(),
		})
	}

	messages, err := lib.Validate(c.Config.Validator, queryParams)
	if err != nil {
		return response.MakeAPIResponseErrorValidation(ctx, messages)
	}

	res := c.CategoryService.FindAll(ctx.UserContext(), queryParams)
	return response.MakeAPIResponseFromService(ctx, res)
}

func (c *CategoryHandler) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	res := c.CategoryService.FindByID(ctx.UserContext(), id)
	return response.MakeAPIResponseFromService(ctx, res)
}

func (c *CategoryHandler) Create(ctx *fiber.Ctx) error {
	request := new(request.CategorySchema)
	if err := ctx.BodyParser(request); err != nil {
		return response.MakeResponseBadRequest(ctx, response.APIResponseOptions[any]{
			Message: exception.ErrBadRequest.Error(),
		})
	}

	messages, err := lib.Validate(c.Config.Validator, request)
	if err != nil {
		return response.MakeAPIResponseErrorValidation(ctx, messages)
	}

	res := c.CategoryService.Create(ctx.UserContext(), request)
	return response.MakeAPIResponseFromService(ctx, res)
}
