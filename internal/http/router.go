package http

import (
	"github.com/bagusyanuar/go-erp/internal/app/di"
	"github.com/bagusyanuar/go-erp/internal/config"
)

func NewRouter(cfg *config.AppConfig, handler *di.HandlerContainer) {
	app := cfg.App
	app.Get("/", handler.Home.Index)
	app.Post("/auth", handler.Auth.Login)
	app.Get("/user", handler.User.FindAll)
	app.Post("/user", handler.User.Create)
	app.Get("/user/:id", handler.User.FindByID)

	app.Post("/unit", handler.Unit.Create)
	app.Get("/unit", handler.Unit.FindAll)
	app.Get("/unit/:id", handler.Unit.FindByID)
}
