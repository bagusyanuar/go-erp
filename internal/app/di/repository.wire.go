package di

import (
	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/internal/domain/repository"
)

type RepositoryContainer struct {
	Auth repository.AuthRepository
	User repository.UserRepository
	Unit repository.UnitRepository
}

func InitRepository(cfg *config.AppConfig) *RepositoryContainer {
	return &RepositoryContainer{
		Auth: repository.NewAuthRepository(cfg.DB),
		User: repository.NewUserRepository(cfg.DB),
		Unit: repository.NewUnitRepository(cfg.DB),
	}
}
