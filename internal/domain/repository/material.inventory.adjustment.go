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
	MaterialInventoryAdjustmentRepository interface {
		Create(ctx context.Context, materialInventoryAdjustment *entity.MaterialInventoryAdjustment) response.RepositoryResponse[any]
		FindAll(ctx context.Context, queryParams *request.MaterialInventoryAdjustmetQuery) response.RepositoryResponse[[]entity.MaterialInventoryAdjustment]
	}

	materialnventoryAdjustmentRepositoryImpl struct {
		DB *gorm.DB
	}
)

func NewMaterialInventoryAdjustmentRepository(db *gorm.DB) MaterialInventoryAdjustmentRepository {
	return &materialnventoryAdjustmentRepositoryImpl{
		DB: db,
	}
}

// Create implements MaterialInventoryAdjustmentRepository.
func (repository *materialnventoryAdjustmentRepositoryImpl) Create(ctx context.Context, materialInventoryAdjustment *entity.MaterialInventoryAdjustment) response.RepositoryResponse[any] {
	tx := repository.DB.WithContext(ctx)
	if err := tx.
		Omit(clause.Associations).
		Create(&materialInventoryAdjustment).Error; err != nil {
		return response.MakeRepositoryError[any](err)
	}
	return response.MakeRepositorySuccess[any](nil, nil)
}

// FindAll implements MaterialInventoryAdjustmentRepository.
func (repository *materialnventoryAdjustmentRepositoryImpl) FindAll(
	ctx context.Context,
	queryParams *request.MaterialInventoryAdjustmetQuery,
) response.RepositoryResponse[[]entity.MaterialInventoryAdjustment] {

	tx := repository.DB.WithContext(ctx)
	defaultQuery := repository.defaultQuery(tx, queryParams)

	var totalRows int64
	if err := defaultQuery.Model(&entity.MaterialInventoryAdjustment{}).
		Count(&totalRows).Error; err != nil {
		return response.MakeRepositoryError[[]entity.MaterialInventoryAdjustment](err)
	}

	var data []entity.MaterialInventoryAdjustment
	if err := defaultQuery.
		Scopes(pagination.Paginate(tx, queryParams.Page, queryParams.PageSize)).
		Find(&data).Error; err != nil {
		return response.MakeRepositoryError[[]entity.MaterialInventoryAdjustment](err)
	}

	meta := pagination.MakeMetaPagination(queryParams.Page, queryParams.PageSize, totalRows)

	return response.MakeRepositorySuccess(data, meta)
}

func (repository *materialnventoryAdjustmentRepositoryImpl) defaultQuery(tx *gorm.DB, queryParams *request.MaterialInventoryAdjustmetQuery) *gorm.DB {
	param := fmt.Sprintf("%%%s%%", queryParams.Param)
	sortFieldMap := map[string]string{
		"name":     "materials.name",
		"quantity": "quantity",
	}
	sort := pagination.GetSortField(queryParams.Sort, "materials.name", sortFieldMap)
	order := pagination.GetOrder(queryParams.Order)

	tx = tx.
		Preload("Material").
		Preload("Unit").
		Preload("Author").
		Joins("JOIN materials ON materials.id = material_inventory_adjustments.material_id").
		Scopes(
			repository.filterByParam(param),
			pagination.SortScope(sort, order),
		).
		Group("material_inventory_adjustments.id, materials.name")

	return tx
}

func (repository *materialnventoryAdjustmentRepositoryImpl) filterByParam(param string) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if param == "" {
			return tx
		}
		return tx.
			Where("materials.name ILIKE ?", param).
			Where("materials.deleted_at IS NULL")
	}
}
