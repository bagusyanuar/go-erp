package di

import (
	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/internal/http/handler"
)

type HandlerContainer struct {
	Home *handler.HomeHandler
	User *handler.UserHandler
}

func InitHandler(cfg *config.AppConfig, serviceContainer *ServiceContainer) *HandlerContainer {
	return &HandlerContainer{
		Home: handler.NewHomeHandler(cfg.Logger),
		User: handler.NewUserHandler(serviceContainer.User, cfg.Validator),
	}
}
