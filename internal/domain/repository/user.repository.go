package repository

import (
	"context"

	"github.com/bagusyanuar/go-erp/internal/domain/entity"
	"github.com/bagusyanuar/go-erp/pkg/lib"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		FindAll(ctx context.Context) lib.RepositoryResponse[[]entity.User]
		FindByID(ctx context.Context, id string) lib.RepositoryResponse[*entity.User]
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

// FindAll implements UserRepository.
func (repository *userRepositoryImpl) FindAll(ctx context.Context) lib.RepositoryResponse[[]entity.User] {
	var data []entity.User
	tx := repository.DB.WithContext(ctx)

	if err := tx.Find(&data).Error; err != nil {
		return lib.MakeRepositoryError[[]entity.User](err)
	}
	return lib.MakeRepositorySuccess(data, nil)
}

// FindByID implements UserRepository.
func (repository *userRepositoryImpl) FindByID(ctx context.Context, id string) lib.RepositoryResponse[*entity.User] {
	var data *entity.User
	tx := repository.DB.WithContext(ctx)
	if err := tx.Where("id = ?", id).
		First(&data).Error; err != nil {
		return lib.MakeRepositoryError[*entity.User](err)
	}
	return lib.MakeRepositorySuccess(data, nil)
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
