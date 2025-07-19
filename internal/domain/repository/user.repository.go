package repository

import (
	"context"

	"github.com/bagusyanuar/go-erp/internal/domain/entity"
	"github.com/bagusyanuar/go-erp/pkg/response"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Create(ctx context.Context, user *entity.User) response.ServiceResponse[any]
	}

	userRepositoryImpl struct {
		DB *gorm.DB
	}
)

// Create implements usecase.UserRepository.
func (repository *userRepositoryImpl) Create(ctx context.Context, user *entity.User) response.ServiceResponse[any] {
	res := response.ServiceResponse[any]{
		Status: response.InternalServerError,
	}

	tx := repository.DB.WithContext(ctx)
	if err := tx.
		Create(&user).Error; err != nil {
		res.Error = err
		return res
	}

	res.Status = response.Created
	return res
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		DB: db,
	}
}
