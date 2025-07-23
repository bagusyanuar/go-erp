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

type AuthHandler struct {
	AuthService service.AuthService
	Config      *config.AppConfig
}

func NewAuthHandler(authService service.AuthService, cfg *config.AppConfig) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
		Config:      cfg,
	}
}

func (c *AuthHandler) Login(ctx *fiber.Ctx) error {
	request := new(request.LoginSchema)
	if err := ctx.BodyParser(request); err != nil {
		return response.MakeResponseBadRequest(ctx, response.APIResponseOptions[any]{
			Message: exception.ErrBadRequest.Error(),
		})
	}

	messages, err := lib.Validate(c.Config.Validator, request)
	if err != nil {
		return response.MakeAPIResponseErrorValidation(ctx, messages)
	}

	res := c.AuthService.Login(ctx.UserContext(), request)
	return response.MakeAPIResponseFromService(ctx, res)
}
