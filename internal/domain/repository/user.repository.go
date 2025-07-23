package repository

import (
	"context"
	"fmt"

	"github.com/bagusyanuar/go-erp/internal/delivery/request"
	"github.com/bagusyanuar/go-erp/internal/domain/entity"
	"github.com/bagusyanuar/go-erp/pkg/lib/pagination"
	"github.com/bagusyanuar/go-erp/pkg/lib/response"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		FindAll(ctx context.Context, queryParams *request.UserQuery) response.RepositoryResponse[[]entity.User]
		FindByID(ctx context.Context, id string) response.RepositoryResponse[*entity.User]
		Create(ctx context.Context, user *entity.User) response.RepositoryResponse[any]
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

func (repository *userRepositoryImpl) defaultQuery(tx *gorm.DB, queryParams *request.UserQuery) *gorm.DB {
	param := fmt.Sprintf("%%%s%%", queryParams.Param)
	tx = tx.Where("email ILIKE ?", param).Where("email != ?", "superdev@web.id")

	return tx
}

// FindAll implements UserRepository.
func (repository *userRepositoryImpl) FindAll(ctx context.Context, queryParams *request.UserQuery) response.RepositoryResponse[[]entity.User] {

	tx := repository.DB.WithContext(ctx)
	defaultQuery := repository.defaultQuery(tx, queryParams)

	var totalRows int64
	if err := defaultQuery.Model(&entity.User{}).
		Count(&totalRows).Error; err != nil {
		return response.MakeRepositoryError[[]entity.User](err)
	}

	var data []entity.User
	if err := defaultQuery.
		Scopes(pagination.Paginate(tx, queryParams.Page, queryParams.PageSize)).
		Find(&data).Error; err != nil {
		return response.MakeRepositoryError[[]entity.User](err)
	}

	meta := pagination.MakeMetaPagination(queryParams.Page, queryParams.PageSize, totalRows)

	return response.MakeRepositorySuccess(data, meta)
}

// FindByID implements UserRepository.
func (repository *userRepositoryImpl) FindByID(ctx context.Context, id string) response.RepositoryResponse[*entity.User] {
	var data *entity.User
	tx := repository.DB.WithContext(ctx)
	if err := tx.Where("id = ?", id).
		Where("email != ?", "superdev@web.id").
		First(&data).Error; err != nil {
		return response.MakeRepositoryError[*entity.User](err)
	}
	return response.MakeRepositorySuccess(data, nil)
}

// Create implements UserRepository.
func (repository *userRepositoryImpl) Create(ctx context.Context, user *entity.User) response.RepositoryResponse[any] {
	tx := repository.DB.WithContext(ctx)
	if err := tx.
		Create(&user).Error; err != nil {
		return response.MakeRepositoryError[any](err)
	}
	return response.MakeRepositorySuccess[any](nil, nil)
}
