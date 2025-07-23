package repository

import (
	"context"

	"github.com/bagusyanuar/go-erp/internal/delivery/request"
	"github.com/bagusyanuar/go-erp/internal/domain/entity"
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

// Create implements MaterialCategoryRepository.
func (repository *materialCategoryRepositoryImpl) Create(ctx context.Context, materialCategory *entity.MaterialCategory) response.RepositoryResponse[any] {
	panic("unimplemented")
}

// FindAll implements MaterialCategoryRepository.
func (repository *materialCategoryRepositoryImpl) FindAll(ctx context.Context, queryParams *request.MaterialCategoryQuery) response.RepositoryResponse[[]entity.MaterialCategory] {
	panic("unimplemented")
}

// FindByID implements MaterialCategoryRepository.
func (repository *materialCategoryRepositoryImpl) FindByID(ctx context.Context, id string) response.RepositoryResponse[*entity.MaterialCategory] {
	panic("unimplemented")
}
