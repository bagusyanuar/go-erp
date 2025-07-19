package handler

import (
	"github.com/bagusyanuar/go-erp/internal/delivery/request"
	"github.com/bagusyanuar/go-erp/internal/service"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (c *UserHandler) Create(ctx *fiber.Ctx) error {

	request := new(request.UserRequest)

	response := c.UserService.Create(ctx.UserContext(), request)
	return ctx.Status(int(response.Status)).JSON(&fiber.Map{
		"status":  response.Status,
		"message": "message",
	})
}
