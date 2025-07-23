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
	UnitRepository interface {
		FindAll(ctx context.Context, queryParams *request.UnitQuery) response.RepositoryResponse[[]entity.Unit]
		FindByID(ctx context.Context, id string) response.RepositoryResponse[*entity.Unit]
		Create(ctx context.Context, unit *entity.Unit) response.RepositoryResponse[any]
	}

	unitRepositoryImpl struct {
		DB *gorm.DB
	}
)

func NewUnitRepository(db *gorm.DB) UnitRepository {
	return &unitRepositoryImpl{
		DB: db,
	}
}

// FindAll implements UnitRepository.
func (repository *unitRepositoryImpl) FindAll(ctx context.Context, queryParams *request.UnitQuery) response.RepositoryResponse[[]entity.Unit] {

	tx := repository.DB.WithContext(ctx)
	defaultQuery := repository.defaultQuery(tx, queryParams)

	var totalRows int64
	if err := defaultQuery.Model(&entity.Unit{}).
		Count(&totalRows).Error; err != nil {
		return response.MakeRepositoryError[[]entity.Unit](err)
	}

	var data []entity.Unit
	if err := defaultQuery.
		Scopes(pagination.Paginate(tx, queryParams.Page, queryParams.PageSize)).
		Find(&data).Error; err != nil {
		return response.MakeRepositoryError[[]entity.Unit](err)
	}

	meta := pagination.MakeMetaPagination(queryParams.Page, queryParams.PageSize, totalRows)

	return response.MakeRepositorySuccess(data, meta)
}

// FindByID implements UnitRepository.
func (repository *unitRepositoryImpl) FindByID(ctx context.Context, id string) response.RepositoryResponse[*entity.Unit] {
	var data *entity.Unit
	tx := repository.DB.WithContext(ctx)
	if err := tx.Where("id = ?", id).
		First(&data).Error; err != nil {
		return response.MakeRepositoryError[*entity.Unit](err)
	}
	return response.MakeRepositorySuccess(data, nil)
}

// Create implements UnitRepository.
func (repository *unitRepositoryImpl) Create(ctx context.Context, unit *entity.Unit) response.RepositoryResponse[any] {
	tx := repository.DB.WithContext(ctx)
	if err := tx.
		Create(&unit).Error; err != nil {
		return response.MakeRepositoryError[any](err)
	}
	return response.MakeRepositorySuccess[any](nil, nil)
}

func (repository *unitRepositoryImpl) defaultQuery(tx *gorm.DB, queryParams *request.UnitQuery) *gorm.DB {
	param := fmt.Sprintf("%%%s%%", queryParams.Param)
	tx = tx.Where("name ILIKE ?", param)

	return tx
}
