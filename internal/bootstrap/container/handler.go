package container

import (
	"github.com/bagusyanuar/go-erp/internal/bootstrap"
	"github.com/bagusyanuar/go-erp/internal/http/handler"
)

type HandlerContainer struct {
	Home *handler.HomeHandler
	User *handler.UserHandler
}

func InitHandler(cfg *bootstrap.AppConfig, service *ServiceContainer) *HandlerContainer {
	return &HandlerContainer{
		Home: handler.NewHomeHandler(cfg.Logger),
		User: handler.NewUserHandler(service.User),
	}
}
