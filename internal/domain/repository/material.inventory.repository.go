package repository

import (
	"context"

	"github.com/bagusyanuar/go-erp/internal/domain/entity"
	"github.com/bagusyanuar/go-erp/pkg/lib/response"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	MaterialInventoryRepository interface {
		Create(ctx context.Context, materialInventory *entity.MaterialInventory) response.RepositoryResponse[any]
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
