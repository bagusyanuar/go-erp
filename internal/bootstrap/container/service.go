package container

import (
	"github.com/bagusyanuar/go-erp/internal/bootstrap"
	"github.com/bagusyanuar/go-erp/internal/service"
	"github.com/bagusyanuar/go-erp/internal/usecase"
)

type ServiceContainer struct {
	User usecase.UserService
}

func InitService(cfg *bootstrap.AppConfig, repository *RepositoryContainer) *ServiceContainer {
	return &ServiceContainer{
		User: service.NewUserService(repository.User),
	}
}
