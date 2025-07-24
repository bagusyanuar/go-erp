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
	CategoryRepository interface {
		FindAll(ctx context.Context, queryParams *request.CategoryQuery) response.RepositoryResponse[[]entity.Category]
		FindByID(ctx context.Context, id string) response.RepositoryResponse[*entity.Category]
		Create(ctx context.Context, category *entity.Category) response.RepositoryResponse[any]
	}

	categoryRepositoryImpl struct {
		DB *gorm.DB
	}
)

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepositoryImpl{
		DB: db,
	}
}

func (repository *categoryRepositoryImpl) defaultQuery(tx *gorm.DB, queryParams *request.CategoryQuery) *gorm.DB {
	param := fmt.Sprintf("%%%s%%", queryParams.Param)
	tx = tx.Where("name ILIKE ?", param)

	return tx
}

// FindAll implements CategoryRepository.
func (repository *categoryRepositoryImpl) FindAll(ctx context.Context, queryParams *request.CategoryQuery) response.RepositoryResponse[[]entity.Category] {
	tx := repository.DB.WithContext(ctx)
	defaultQuery := repository.defaultQuery(tx, queryParams)

	var totalRows int64
	if err := defaultQuery.Model(&entity.Category{}).
		Count(&totalRows).Error; err != nil {
		return response.MakeRepositoryError[[]entity.Category](err)
	}

	var data []entity.Category
	if err := defaultQuery.
		Scopes(pagination.Paginate(tx, queryParams.Page, queryParams.PageSize)).
		Find(&data).Error; err != nil {
		return response.MakeRepositoryError[[]entity.Category](err)
	}

	meta := pagination.MakeMetaPagination(queryParams.Page, queryParams.PageSize, totalRows)

	return response.MakeRepositorySuccess(data, meta)
}

// FindByID implements CategoryRepository.
func (repository *categoryRepositoryImpl) FindByID(ctx context.Context, id string) response.RepositoryResponse[*entity.Category] {
	var data *entity.Category
	tx := repository.DB.WithContext(ctx)
	if err := tx.Where("id = ?", id).
		First(&data).Error; err != nil {
		return response.MakeRepositoryError[*entity.Category](err)
	}
	return response.MakeRepositorySuccess(data, nil)
}

// Create implements CategoryRepository.
func (repository *categoryRepositoryImpl) Create(ctx context.Context, category *entity.Category) response.RepositoryResponse[any] {
	tx := repository.DB.WithContext(ctx)
	if err := tx.
		Create(&category).Error; err != nil {
		return response.MakeRepositoryError[any](err)
	}
	return response.MakeRepositorySuccess[any](nil, nil)
}
