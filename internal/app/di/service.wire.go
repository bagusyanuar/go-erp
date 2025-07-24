package di

import (
	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/internal/service"
)

type ServiceContainer struct {
	Auth             service.AuthService
	User             service.UserService
	Unit             service.UnitService
	MaterialCategory service.MaterialCategoryService
}

func InitService(cfg *config.AppConfig, repositoryContainer *RepositoryContainer) *ServiceContainer {
	return &ServiceContainer{
		Auth:             service.NewAuthService(repositoryContainer.Auth, cfg),
		User:             service.NewUserService(repositoryContainer.User),
		Unit:             service.NewUnitService(repositoryContainer.Unit),
		MaterialCategory: service.NewMaterialCategoryService(repositoryContainer.MaterialCategory),
	}
}
