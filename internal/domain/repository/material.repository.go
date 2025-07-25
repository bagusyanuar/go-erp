package repository

import (
	"context"
	"fmt"

	"github.com/bagusyanuar/go-erp/internal/delivery/request"
	"github.com/bagusyanuar/go-erp/internal/domain/entity"
	"github.com/bagusyanuar/go-erp/pkg/lib/pagination"
	"github.com/bagusyanuar/go-erp/pkg/lib/response"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	MaterialRepository interface {
		FindAll(ctx context.Context, queryParams *request.MaterialQuery) response.RepositoryResponse[[]entity.Material]
		FindByID(ctx context.Context, id string) response.RepositoryResponse[*entity.Material]
		Create(ctx context.Context, material *entity.Material, categories []string) response.RepositoryResponse[any]
	}

	materialRepositoryImpl struct {
		DB *gorm.DB
	}
)

func NewMaterialRepository(db *gorm.DB) MaterialRepository {
	return &materialRepositoryImpl{
		DB: db,
	}
}

// Create implements MaterialRepository.
func (repository *materialRepositoryImpl) Create(ctx context.Context, material *entity.Material, categories []string) response.RepositoryResponse[any] {
	tx := repository.DB.WithContext(ctx)

	err := tx.Transaction(func(tx *gorm.DB) error {
		var dataCategories []entity.Category
		if err := tx.Where("id IN ?", categories).Find(&dataCategories).Error; err != nil {
			return err
		}

		if err := tx.Omit(clause.Associations).Create(&material).Error; err != nil {
			return err
		}

		if err := tx.Model(&material).Association("Categories").Replace(dataCategories); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return response.MakeRepositoryError[any](err)
	}
	return response.MakeRepositorySuccess[any](nil, nil)
}

// FindAll implements MaterialRepository.
func (repository *materialRepositoryImpl) FindAll(ctx context.Context, queryParams *request.MaterialQuery) response.RepositoryResponse[[]entity.Material] {
	tx := repository.DB.WithContext(ctx)
	defaultQuery := repository.defaultQuery(tx, queryParams)

	var totalRows int64
	if err := defaultQuery.Model(&entity.Material{}).
		Count(&totalRows).Error; err != nil {
		return response.MakeRepositoryError[[]entity.Material](err)
	}

	var data []entity.Material
	if err := defaultQuery.
		Preload("Categories").
		Scopes(pagination.Paginate(tx, queryParams.Page, queryParams.PageSize)).
		Find(&data).Error; err != nil {
		return response.MakeRepositoryError[[]entity.Material](err)
	}

	meta := pagination.MakeMetaPagination(queryParams.Page, queryParams.PageSize, totalRows)

	return response.MakeRepositorySuccess(data, meta)
}

// FindByID implements MaterialRepository.
func (repository *materialRepositoryImpl) FindByID(ctx context.Context, id string) response.RepositoryResponse[*entity.Material] {
	var data *entity.Material
	tx := repository.DB.WithContext(ctx)
	if err := tx.Where("id = ?", id).
		Preload("Categories").
		First(&data).Error; err != nil {
		return response.MakeRepositoryError[*entity.Material](err)
	}
	return response.MakeRepositorySuccess(data, nil)
}

func (repository *materialRepositoryImpl) defaultQuery(tx *gorm.DB, queryParams *request.MaterialQuery) *gorm.DB {
	param := fmt.Sprintf("%%%s%%", queryParams.Param)
	tx = tx.Where("name ILIKE ?", param)

	return tx
}
