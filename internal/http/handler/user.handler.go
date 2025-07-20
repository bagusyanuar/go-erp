package handler

import (
	"github.com/bagusyanuar/go-erp/internal/delivery/request"
	"github.com/bagusyanuar/go-erp/internal/domain/dto"
	"github.com/bagusyanuar/go-erp/internal/pkg/helper"
	"github.com/bagusyanuar/go-erp/internal/pkg/myexception"
	"github.com/bagusyanuar/go-erp/internal/service"
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

func (c *UserHandler) Create(ctx *fiber.Ctx) error {
	request := new(request.UserRequest)

	if err := ctx.BodyParser(request); err != nil {
		return helper.MakeResponse(ctx, dto.APIResponse[any]{
			Code:    dto.BadRequest,
			Message: myexception.ErrBadRequest.Error(),
		})
	}

	messages, err := helper.Validate(c.Validator, request)
	if err != nil {
		return helper.MakeResponse(ctx, dto.APIResponse[any]{
			Code:    dto.UnprocessableEntity,
			Message: myexception.ErrUnprocessableEntity.Error(), // custom formatter,
			Data:    messages,
		})
	}

	response := c.UserService.Create(ctx.UserContext(), request)
	return helper.MakeResponse(ctx, dto.FromService(response, nil))
}
