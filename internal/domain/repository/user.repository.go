package repository

import (
	"context"

	"github.com/bagusyanuar/go-erp/internal/domain/dto"
	"github.com/bagusyanuar/go-erp/internal/domain/entity"
	"github.com/bagusyanuar/go-erp/internal/pkg/myexception"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Create(ctx context.Context, user *entity.User) dto.ServiceResponse[any]
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
func (repository *userRepositoryImpl) Create(ctx context.Context, user *entity.User) dto.ServiceResponse[any] {
	res := dto.ServiceResponse[any]{
		Status:  dto.InternalServerError,
		Message: "internal server error",
		Error:   myexception.ErrUnknown,
	}

	tx := repository.DB.WithContext(ctx)
	if err := tx.
		Create(&user).Error; err != nil {
		res.Error = err
		res.Message = err.Error()
		return res
	}

	res.Status = dto.Created
	res.Message = "successfully create user"
	return res
}
