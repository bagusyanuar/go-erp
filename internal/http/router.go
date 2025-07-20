package http

import (
	"github.com/bagusyanuar/go-erp/internal/app/di"
	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(cfg *config.AppConfig, handler *di.HandlerContainer) *fiber.App {
	app := fiber.New()
	app.Get("/", handler.Home.Index)
	app.Get("/user", handler.User.FindAll)
	app.Post("/user", handler.User.Create)
	app.Get("/user/:id", handler.User.FindByID)
	return app
}
