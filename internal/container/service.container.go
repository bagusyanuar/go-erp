package container

import (
	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/internal/service"
)

type ServiceContainer struct {
	User service.UserService
}

func InitService(cfg *config.AppConfig, repositoryContainer *RepositoryContainer) *ServiceContainer {
	return &ServiceContainer{
		User: service.NewUserService(repositoryContainer.User),
	}
}
