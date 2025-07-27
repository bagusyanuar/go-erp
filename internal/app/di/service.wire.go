package di

import (
	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/internal/service"
)

type ServiceContainer struct {
	Auth                        service.AuthService
	User                        service.UserService
	Unit                        service.UnitService
	Category                    service.CategoryService
	Material                    service.MaterialService
	MaterialInventory           service.MaterialInventoryService
	MaterialInventoryAdjustment service.MaterialInventoryAdjustmentService
}

func InitService(cfg *config.AppConfig, repositoryContainer *RepositoryContainer) *ServiceContainer {
	return &ServiceContainer{
		Auth:                        service.NewAuthService(repositoryContainer.Auth, cfg),
		User:                        service.NewUserService(repositoryContainer.User),
		Unit:                        service.NewUnitService(repositoryContainer.Unit),
		Category:                    service.NewCategoryService(repositoryContainer.Category),
		Material:                    service.NewMaterialService(repositoryContainer.Material),
		MaterialInventory:           service.NewMaterialInventoryService(repositoryContainer.MaterialInventory),
		MaterialInventoryAdjustment: service.NewMaterialInventoryAdjustmentService(repositoryContainer.MaterialInventoryAdjustment),
	}
}
