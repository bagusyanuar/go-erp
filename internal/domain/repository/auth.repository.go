package repository

import (
	"context"

	"github.com/bagusyanuar/go-erp/internal/domain/entity"
	"github.com/bagusyanuar/go-erp/pkg/lib/response"
	"gorm.io/gorm"
)

type (
	AuthRepository interface {
		Login(ctx context.Context, email string) response.RepositoryResponse[*entity.User]
	}

	authRepsitoryImpl struct {
		DB *gorm.DB
	}
)

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepsitoryImpl{
		DB: db,
	}
}

// Login implements AuthRepository.
func (repository *authRepsitoryImpl) Login(ctx context.Context, email string) response.RepositoryResponse[*entity.User] {
	var data *entity.User
	tx := repository.DB.WithContext(ctx)
	if err := tx.Where("email = ?", email).
		First(&data).Error; err != nil {
		return response.MakeRepositoryError[*entity.User](err)
	}
	return response.MakeRepositorySuccess(data, nil)
}
