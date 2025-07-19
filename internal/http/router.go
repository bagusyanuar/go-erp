package http

import (
	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/internal/container"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(cfg *config.AppConfig, handler *container.HandlerContainer) *fiber.App {
	app := fiber.New()
	app.Get("/", handler.Home.Index)
	app.Get("/user", handler.User.Create)
	return app
}
