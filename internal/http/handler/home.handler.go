package handler

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type HomeHandler struct {
	log *zap.Logger
}

func NewHomeHandler(log *zap.Logger) *HomeHandler {
	return &HomeHandler{
		log: log,
	}
}

func (c *HomeHandler) Index(ctx *fiber.Ctx) error {
	c.log.Info("server running well..")
	return ctx.Status(200).JSON(&fiber.Map{
		"app_name":    "go-erp-backend",
		"app_version": "v1.0",
	})
}
