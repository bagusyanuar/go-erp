package http

import (
	"github.com/bagusyanuar/go-erp/internal/app/di"
	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(cfg *config.AppConfig, handler *di.HandlerContainer) *fiber.App {
	app := fiber.New()
	app.Get("/", handler.Home.Index)
	app.Post("/user", handler.User.Create)
	return app
}
