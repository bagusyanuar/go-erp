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

type UserHandler struct {
	UserService service.UserService
	Config      *config.AppConfig
}

func NewUserHandler(userService service.UserService, cfg *config.AppConfig) *UserHandler {
	return &UserHandler{
		UserService: userService,
		Config:      cfg,
	}
}

func (c *UserHandler) FindAll(ctx *fiber.Ctx) error {
	queryParams := new(request.UserQuery)
	if err := ctx.QueryParser(queryParams); err != nil {
		return response.MakeResponseBadRequest(ctx, response.APIResponseOptions[any]{
			Message: exception.ErrBadRequest.Error(),
		})
	}

	messages, err := lib.Validate(c.Config.Validator, queryParams)
	if err != nil {
		return response.MakeAPIResponseErrorValidation(ctx, messages)
	}

	res := c.UserService.FindAll(ctx.UserContext(), queryParams)
	return response.MakeAPIResponseFromService(ctx, res)
}

func (c *UserHandler) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	res := c.UserService.FindByID(ctx.UserContext(), id)
	return response.MakeAPIResponseFromService(ctx, res)
}

func (c *UserHandler) Create(ctx *fiber.Ctx) error {
	request := new(request.UserSchema)
	if err := ctx.BodyParser(request); err != nil {
		return response.MakeResponseBadRequest(ctx, response.APIResponseOptions[any]{
			Message: exception.ErrBadRequest.Error(),
		})
	}

	messages, err := lib.Validate(c.Config.Validator, request)
	if err != nil {
		return response.MakeAPIResponseErrorValidation(ctx, messages)
	}

	res := c.UserService.Create(ctx.UserContext(), request)
	return response.MakeAPIResponseFromService(ctx, res)
}
