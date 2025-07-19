package container

import (
	"github.com/bagusyanuar/go-erp/internal/bootstrap"
	"github.com/bagusyanuar/go-erp/internal/domain/repository"
)

type RepositoryContainer struct {
	User repository.UserRepository
}

func InitRepository(cfg *bootstrap.AppConfig) *RepositoryContainer {
	return &RepositoryContainer{
		User: repository.NewUserRepository(cfg.DB),
	}
}
