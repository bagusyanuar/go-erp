package http

import (
	"github.com/bagusyanuar/go-erp/internal/app/di"
	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/internal/http/middleware"
)

func NewRouter(cfg *config.AppConfig, handler *di.HandlerContainer) {
	app := cfg.App

	jwtMiddleware := middleware.VerifyJWT(cfg)
	app.Get("/", handler.Home.Index)
	app.Post("/login", handler.Auth.Login)

	user := app.Group("/user")
	user.Get("/", handler.User.FindAll)
	user.Post("/", handler.User.Create)
	user.Get("/:id", handler.User.FindByID)

	unit := app.Group("/unit")
	unit.Post("/", handler.Unit.Create)
	unit.Get("/", jwtMiddleware, handler.Unit.FindAll)
	unit.Get("/:id", handler.Unit.FindByID)

	materialCategory := app.Group("/material-category")
	materialCategory.Post("/", handler.MaterialCategory.Create)
	materialCategory.Get("/", handler.MaterialCategory.FindAll)
	materialCategory.Get("/:id", handler.MaterialCategory.FindByID)
}
