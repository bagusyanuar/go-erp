package repository

import (
	"context"

	"github.com/bagusyanuar/go-erp/internal/delivery/request"
	"github.com/bagusyanuar/go-erp/internal/domain/entity"
	"github.com/bagusyanuar/go-erp/pkg/lib"
	"gorm.io/gorm"
)

type (
	UnitRepository interface {
		FindAll(ctx context.Context, queryParams *request.UnitQuery) lib.RepositoryResponse[[]entity.Unit]
		FindByID(ctx context.Context, id string) lib.RepositoryResponse[*entity.Unit]
		Create(ctx context.Context, unit *entity.Unit) lib.RepositoryResponse[any]
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

// Create implements UnitRepository.
func (repository *unitRepositoryImpl) Create(ctx context.Context, unit *entity.Unit) lib.RepositoryResponse[any] {
	tx := repository.DB.WithContext(ctx)
	if err := tx.
		Create(&unit).Error; err != nil {
		return lib.MakeRepositoryError[any](err)
	}
	return lib.MakeRepositorySuccess[any](nil, nil)
}

// FindAll implements UnitRepository.
func (repository *unitRepositoryImpl) FindAll(ctx context.Context, queryParams *request.UnitQuery) lib.RepositoryResponse[[]entity.Unit] {
	var data []entity.Unit
	tx := repository.DB.WithContext(ctx)

	if err := tx.Find(&data).Error; err != nil {
		return lib.MakeRepositoryError[[]entity.Unit](err)
	}
	return lib.MakeRepositorySuccess(data, nil)
}

// FindByID implements UnitRepository.
func (repository *unitRepositoryImpl) FindByID(ctx context.Context, id string) lib.RepositoryResponse[*entity.Unit] {
	var data *entity.Unit
	tx := repository.DB.WithContext(ctx)
	if err := tx.Where("id = ?", id).
		First(&data).Error; err != nil {
		return lib.MakeRepositoryError[*entity.Unit](err)
	}
	return lib.MakeRepositorySuccess(data, nil)
}
