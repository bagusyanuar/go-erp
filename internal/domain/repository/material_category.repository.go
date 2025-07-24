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
	MaterialCategoryRepository interface {
		FindAll(ctx context.Context, queryParams *request.MaterialCategoryQuery) response.RepositoryResponse[[]entity.MaterialCategory]
		FindByID(ctx context.Context, id string) response.RepositoryResponse[*entity.MaterialCategory]
		Create(ctx context.Context, materialCategory *entity.MaterialCategory) response.RepositoryResponse[any]
	}

	materialCategoryRepositoryImpl struct {
		DB *gorm.DB
	}
)

func NewMaterialCategoryRepository(db *gorm.DB) MaterialCategoryRepository {
	return &materialCategoryRepositoryImpl{
		DB: db,
	}
}

// FindAll implements MaterialCategoryRepository.
func (repository *materialCategoryRepositoryImpl) FindAll(ctx context.Context, queryParams *request.MaterialCategoryQuery) response.RepositoryResponse[[]entity.MaterialCategory] {
	tx := repository.DB.WithContext(ctx)
	defaultQuery := repository.defaultQuery(tx, queryParams)

	var totalRows int64
	if err := defaultQuery.Model(&entity.MaterialCategory{}).
		Count(&totalRows).Error; err != nil {
		return response.MakeRepositoryError[[]entity.MaterialCategory](err)
	}

	var data []entity.MaterialCategory
	if err := defaultQuery.
		Scopes(pagination.Paginate(tx, queryParams.Page, queryParams.PageSize)).
		Find(&data).Error; err != nil {
		return response.MakeRepositoryError[[]entity.MaterialCategory](err)
	}

	meta := pagination.MakeMetaPagination(queryParams.Page, queryParams.PageSize, totalRows)

	return response.MakeRepositorySuccess(data, meta)
}

// FindByID implements MaterialCategoryRepository.
func (repository *materialCategoryRepositoryImpl) FindByID(ctx context.Context, id string) response.RepositoryResponse[*entity.MaterialCategory] {
	var data *entity.MaterialCategory
	tx := repository.DB.WithContext(ctx)
	if err := tx.Where("id = ?", id).
		First(&data).Error; err != nil {
		return response.MakeRepositoryError[*entity.MaterialCategory](err)
	}
	return response.MakeRepositorySuccess(data, nil)
}

// Create implements MaterialCategoryRepository.
func (repository *materialCategoryRepositoryImpl) Create(ctx context.Context, materialCategory *entity.MaterialCategory) response.RepositoryResponse[any] {
	tx := repository.DB.WithContext(ctx)
	if err := tx.
		Create(&materialCategory).Error; err != nil {
		return response.MakeRepositoryError[any](err)
	}
	return response.MakeRepositorySuccess[any](nil, nil)
}

func (repository *materialCategoryRepositoryImpl) defaultQuery(tx *gorm.DB, queryParams *request.MaterialCategoryQuery) *gorm.DB {
	param := fmt.Sprintf("%%%s%%", queryParams.Param)
	tx = tx.Where("name ILIKE ?", param)

	return tx
}
