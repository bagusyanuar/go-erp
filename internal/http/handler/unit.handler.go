package handler

import (
	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/internal/delivery/request"
	"github.com/bagusyanuar/go-erp/internal/service"
	"github.com/bagusyanuar/go-erp/pkg/lib"
	"github.com/bagusyanuar/go-erp/pkg/myexception"
	"github.com/gofiber/fiber/v2"
)

type UnitHandler struct {
	UnitService service.UnitService
	Config      *config.AppConfig
}

func NewUnitHandler(unitService service.UnitService, cfg *config.AppConfig) *UnitHandler {
	return &UnitHandler{
		UnitService: unitService,
		Config:      cfg,
	}
}

func (c *UnitHandler) FindAll(ctx *fiber.Ctx) error {

	queryParams := new(request.UnitQuery)
	if err := ctx.QueryParser(queryParams); err != nil {
		return lib.ResponseBadRequest(ctx, lib.ResponseOptions[any]{
			Message: myexception.ErrBadRequest.Error(),
		})
	}

	messages, err := lib.Validate(c.Config.Validator, queryParams)
	if err != nil {
		return lib.ResponseErrValidation(ctx, messages)
	}

	response := c.UnitService.FindAll(ctx.UserContext(), queryParams)
	return lib.MakeResponse(ctx, lib.FromService(response))
}

func (c *UnitHandler) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	response := c.UnitService.FindByID(ctx.UserContext(), id)
	return lib.MakeResponse(ctx, lib.FromService(response))
}

func (c *UnitHandler) Create(ctx *fiber.Ctx) error {
	request := new(request.UnitSchema)

	if err := ctx.BodyParser(request); err != nil {
		return lib.ResponseBadRequest(ctx, lib.ResponseOptions[any]{
			Message: myexception.ErrBadRequest.Error(),
		})
	}

	messages, err := lib.Validate(c.Config.Validator, request)
	if err != nil {
		return lib.ResponseErrValidation(ctx, messages)
	}

	response := c.UnitService.Create(ctx.UserContext(), request)
	return lib.MakeResponseFromService(ctx, response)
}
