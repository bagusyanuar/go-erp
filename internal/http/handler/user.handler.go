package handler

import (
	"github.com/bagusyanuar/go-erp/internal/delivery/request"
	"github.com/bagusyanuar/go-erp/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService usecase.UserService
}

func NewUserHandler(userService usecase.UserService) *UserHandler {
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
