package repository

import (
	"context"

	"github.com/bagusyanuar/go-erp/internal/domain/entity"
	"github.com/bagusyanuar/go-erp/pkg/lib"
	"gorm.io/gorm"
)

type (
	AuthRepository interface {
		Login(ctx context.Context, email string) lib.RepositoryResponse[*entity.User]
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
func (repository *authRepsitoryImpl) Login(ctx context.Context, email string) lib.RepositoryResponse[*entity.User] {
	var data *entity.User
	tx := repository.DB.WithContext(ctx)
	if err := tx.Where("email = ?", email).
		First(&data).Error; err != nil {
		return lib.MakeRepositoryError[*entity.User](err)
	}
	return lib.MakeRepositorySuccess(data, nil)
}
