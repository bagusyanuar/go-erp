package container

import (
	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/internal/domain/repository"
)

type RepositoryContainer struct {
	User repository.UserRepository
}

func InitRepository(cfg *config.AppConfig) *RepositoryContainer {
	return &RepositoryContainer{
		User: repository.NewUserRepository(cfg.DB),
	}
}
