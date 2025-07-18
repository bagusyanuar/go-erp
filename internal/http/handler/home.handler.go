package handler

import "github.com/gofiber/fiber/v2"

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (c *HomeHandler) Index(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(&fiber.Map{
		"app_name":    "go-erp-backend",
		"app_version": "v1.0",
	})
}
