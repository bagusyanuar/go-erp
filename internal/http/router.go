package http

import (
	"github.com/bagusyanuar/go-erp/internal/bootstrap"
	"github.com/bagusyanuar/go-erp/internal/bootstrap/container"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(cfg *bootstrap.AppConfig, handler *container.HandlerContainer) *fiber.App {
	app := fiber.New()
	app.Get("/", handler.Home.Index)
	app.Get("/user", handler.User.Create)
	return app
}
