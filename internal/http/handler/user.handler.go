package handler

import (
	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/internal/delivery/request"
	"github.com/bagusyanuar/go-erp/internal/service"
	"github.com/bagusyanuar/go-erp/pkg/lib"
	"github.com/bagusyanuar/go-erp/pkg/myexception"
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
		return lib.ResponseBadRequest(ctx, lib.ResponseOptions[any]{
			Message: myexception.ErrBadRequest.Error(),
		})
	}

	messages, err := lib.Validate(c.Config.Validator, queryParams)
	if err != nil {
		return lib.ResponseErrValidation(ctx, messages)
	}

	response := c.UserService.FindAll(ctx.UserContext(), queryParams)
	return lib.MakeResponse(ctx, lib.FromService(response))
}

func (c *UserHandler) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	response := c.UserService.FinByID(ctx.UserContext(), id)
	return lib.MakeResponse(ctx, lib.FromService(response))
}

func (c *UserHandler) Create(ctx *fiber.Ctx) error {
	request := new(request.UserSchema)

	if err := ctx.BodyParser(request); err != nil {
		return lib.ResponseBadRequest(ctx, lib.ResponseOptions[any]{
			Message: myexception.ErrBadRequest.Error(),
		})
	}

	messages, err := lib.Validate(c.Config.Validator, request)
	if err != nil {
		return lib.ResponseErrValidation(ctx, messages)
	}

	response := c.UserService.Create(ctx.UserContext(), request)
	return lib.MakeResponseFromService(ctx, response)
}
