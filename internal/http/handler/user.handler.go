package handler

import (
	"github.com/bagusyanuar/go-erp/internal/delivery/request"
	"github.com/bagusyanuar/go-erp/internal/service"
	"github.com/bagusyanuar/go-erp/pkg/lib"
	"github.com/bagusyanuar/go-erp/pkg/myexception"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService service.UserService
	Validator   *validator.Validate
}

func NewUserHandler(userService service.UserService, validator *validator.Validate) *UserHandler {
	return &UserHandler{
		UserService: userService,
		Validator:   validator,
	}
}

func (c *UserHandler) FindAll(ctx *fiber.Ctx) error {
	response := c.UserService.FindAll(ctx.UserContext())
	return lib.MakeResponse(ctx, lib.FromService(response, nil))
}

func (c *UserHandler) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	response := c.UserService.FinByID(ctx.UserContext(), id)
	return lib.MakeResponse(ctx, lib.FromService(response, nil))
}

func (c *UserHandler) Create(ctx *fiber.Ctx) error {
	request := new(request.UserRequest)

	if err := ctx.BodyParser(request); err != nil {
		return lib.ResponseBadRequest(ctx, lib.ResponseOptions[any]{
			Message: myexception.ErrBadRequest.Error(),
		})
	}

	messages, err := lib.Validate(c.Validator, request)
	if err != nil {
		return lib.ResponseUnproccesableEntity(ctx, lib.ResponseOptions[map[string][]string]{
			Message: myexception.ErrUnprocessableEntity.Error(),
			Data:    messages,
		})
	}

	response := c.UserService.Create(ctx.UserContext(), request)
	return lib.MakeResponse(ctx, lib.FromService(response, nil))
}
