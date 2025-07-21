package handler

import (
	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/internal/delivery/request"
	"github.com/bagusyanuar/go-erp/internal/service"
	"github.com/bagusyanuar/go-erp/pkg/lib"
	"github.com/bagusyanuar/go-erp/pkg/myexception"
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
		return lib.ResponseBadRequest(ctx, lib.ResponseOptions[any]{
			Message: myexception.ErrBadRequest.Error(),
		})
	}

	messages, err := lib.Validate(c.Config.Validator, request)
	if err != nil {
		return lib.ResponseErrValidation(ctx, messages)
	}

	response := c.AuthService.Login(ctx.UserContext(), request)
	return lib.MakeResponseFromService(ctx, response)
}
