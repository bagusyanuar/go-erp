package di

import (
	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/internal/http/handler"
)

type HandlerContainer struct {
	Home *handler.HomeHandler
	Auth *handler.AuthHandler
	User *handler.UserHandler
	Unit *handler.UnitHandler
}

func InitHandler(cfg *config.AppConfig, serviceContainer *ServiceContainer) *HandlerContainer {
	return &HandlerContainer{
		Home: handler.NewHomeHandler(cfg.Logger),
		Auth: handler.NewAuthHandler(serviceContainer.Auth, cfg),
		User: handler.NewUserHandler(serviceContainer.User, cfg),
		Unit: handler.NewUnitHandler(serviceContainer.Unit, cfg),
	}
}
