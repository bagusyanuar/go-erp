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
	MaterialInventoryRepository interface {
		Create(ctx context.Context, materialInventory *entity.MaterialInventory) response.RepositoryResponse[any]
		FindAll(ctx context.Context, queryParams *request.MaterialInventoryQuery) response.RepositoryResponse[[]entity.MaterialInventory]
		FindByID(ctx context.Context, id string) response.RepositoryResponse[*entity.MaterialInventory]
	}

	materialInventoryRepositoryImpl struct {
		DB *gorm.DB
	}
)

func NewMaterialInventoryRepository(db *gorm.DB) MaterialInventoryRepository {
	return &materialInventoryRepositoryImpl{
		DB: db,
	}
}

// Create implements MaterialInventoryRepository.
func (repository *materialInventoryRepositoryImpl) Create(ctx context.Context, materialInventory *entity.MaterialInventory) response.RepositoryResponse[any] {
	tx := repository.DB.WithContext(ctx)
	if err := tx.
		Omit(clause.Associations).
		Create(&materialInventory).Error; err != nil {
		return response.MakeRepositoryError[any](err)
	}
	return response.MakeRepositorySuccess[any](nil, nil)
}

// FindAll implements MaterialInventoryRepository.
func (repository *materialInventoryRepositoryImpl) FindAll(ctx context.Context, queryParams *request.MaterialInventoryQuery) response.RepositoryResponse[[]entity.MaterialInventory] {
	tx := repository.DB.WithContext(ctx)
	defaultQuery := repository.defaultQuery(tx, queryParams)

	var totalRows int64
	if err := defaultQuery.Model(&entity.MaterialInventory{}).
		Count(&totalRows).Error; err != nil {
		return response.MakeRepositoryError[[]entity.MaterialInventory](err)
	}

	var data []entity.MaterialInventory
	if err := defaultQuery.
		Scopes(pagination.Paginate(tx, queryParams.Page, queryParams.PageSize)).
		Find(&data).Error; err != nil {
		return response.MakeRepositoryError[[]entity.MaterialInventory](err)
	}

	meta := pagination.MakeMetaPagination(queryParams.Page, queryParams.PageSize, totalRows)

	return response.MakeRepositorySuccess(data, meta)
}

// FindByID implements MaterialInventoryRepository.
func (repository *materialInventoryRepositoryImpl) FindByID(ctx context.Context, id string) response.RepositoryResponse[*entity.MaterialInventory] {
	var data *entity.MaterialInventory
	tx := repository.DB.WithContext(ctx)
	if err := tx.Where("id = ?", id).
		Preload("Material").
		Preload("Unit").
		Preload("Modificator").
		First(&data).Error; err != nil {
		return response.MakeRepositoryError[*entity.MaterialInventory](err)
	}
	return response.MakeRepositorySuccess(data, nil)
}

func (repository *materialInventoryRepositoryImpl) defaultQuery(tx *gorm.DB, queryParams *request.MaterialInventoryQuery) *gorm.DB {
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
		Preload("Modificator").
		Joins("JOIN materials ON materials.id = material_inventories.material_id").
		Scopes(
			repository.filterByParam(param),
			pagination.SortScope(sort, order),
		).
		Group("material_inventories.id, materials.name")

	return tx
}

func (repository *materialInventoryRepositoryImpl) filterByParam(param string) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if param == "" {
			return tx
		}
		return tx.
			Where("materials.name ILIKE ?", param).
			Where("materials.deleted_at IS NULL")
	}
}
