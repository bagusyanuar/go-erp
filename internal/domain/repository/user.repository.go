package repository

import (
	"context"

	"github.com/bagusyanuar/go-erp/internal/domain/entity"
	"github.com/bagusyanuar/go-erp/pkg/lib"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Create(ctx context.Context, user *entity.User) lib.RepositoryResponse[any]
	}

	userRepositoryImpl struct {
		DB *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		DB: db,
	}
}

// Create implements UserRepository.
func (repository *userRepositoryImpl) Create(ctx context.Context, user *entity.User) lib.RepositoryResponse[any] {
	tx := repository.DB.WithContext(ctx)
	if err := tx.
		Create(&user).Error; err != nil {
		return lib.MakeRepositoryError[any](err)
	}
	return lib.MakeRepositorySuccess[any](nil, nil)
}
