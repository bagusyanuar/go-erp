package container

import (
	"github.com/bagusyanuar/go-erp/internal/bootstrap"
	"github.com/bagusyanuar/go-erp/internal/service"
)

type ServiceContainer struct {
	User service.UserService
}

func InitService(cfg *bootstrap.AppConfig, repository *RepositoryContainer) *ServiceContainer {
	return &ServiceContainer{
		User: service.NewUserService(repository.User),
	}
}
