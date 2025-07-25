package repository

import (
	"context"

	"github.com/bagusyanuar/go-erp/internal/delivery/request"
	"github.com/bagusyanuar/go-erp/internal/domain/entity"
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
	panic("unimplemented")
}

// FindByID implements MaterialRepository.
func (repository *materialRepositoryImpl) FindByID(ctx context.Context, id string) response.RepositoryResponse[*entity.Material] {
	panic("unimplemented")
}
