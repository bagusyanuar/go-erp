package container

import (
	"github.com/bagusyanuar/go-erp/internal/bootstrap"
	"github.com/bagusyanuar/go-erp/internal/domain/repository"
	"github.com/bagusyanuar/go-erp/internal/usecase"
)

type RepositoryContainer struct {
	User usecase.UserRepository
}

func InitRepository(cfg *bootstrap.AppConfig) *RepositoryContainer {
	return &RepositoryContainer{
		User: repository.NewUserRepository(cfg.DB),
	}
}
