package service

import (
	"context"
	"errors"

	"github.com/bagusyanuar/go-erp/internal/delivery/request"
	"github.com/bagusyanuar/go-erp/internal/domain/dto"
	"github.com/bagusyanuar/go-erp/internal/domain/entity"
	"github.com/bagusyanuar/go-erp/internal/domain/repository"
	"github.com/bagusyanuar/go-erp/pkg/lib/response"
	"gorm.io/gorm"
)

type (
	UnitService interface {
		FindAll(ctx context.Context, queryParams *request.UnitQuery) response.ServiceResponse[*[]dto.UnitDTO]
		FindByID(ctx context.Context, id string) response.ServiceResponse[*dto.UnitDTO]
		Create(ctx context.Context, schema *request.UnitSchema) response.ServiceResponse[any]
	}

	unitServiceImpl struct {
		UnitRepository repository.UnitRepository
	}
)

func NewUnitService(unitRepository repository.UnitRepository) UnitService {
	return &unitServiceImpl{
		UnitRepository: unitRepository,
	}
}

// Create implements UnitService.
func (service *unitServiceImpl) Create(ctx context.Context, schema *request.UnitSchema) response.ServiceResponse[any] {
	name := schema.Name

	data := &entity.Unit{
		Name: name,
	}
	repositoryResponse := service.UnitRepository.Create(ctx, data)
	if repositoryResponse.Error != nil {
		return response.ServiceInternalServerError(response.ServiceResponseOptions[any]{
			Error:   repositoryResponse.Error,
			Message: repositoryResponse.Message,
		})
	}
	return response.ServiceCreated(response.ServiceResponseOptions[any]{
		Message: "successfully create unit",
	})
}

// FindAll implements UnitService.
func (service *unitServiceImpl) FindAll(ctx context.Context, queryParams *request.UnitQuery) response.ServiceResponse[*[]dto.UnitDTO] {
	repositoryResponse := service.UnitRepository.FindAll(ctx, queryParams)
	if repositoryResponse.Error != nil {
		return response.ServiceInternalServerError(response.ServiceResponseOptions[*[]dto.UnitDTO]{
			Error:   repositoryResponse.Error,
			Message: repositoryResponse.Message,
		})
	}
	data := dto.ToUnits(repositoryResponse.Data)
	return response.ServiceOK(response.ServiceResponseOptions[*[]dto.UnitDTO]{
		Message: "successfully get units",
		Data:    &data,
		Meta:    repositoryResponse.Meta,
	})
}

// FindByID implements UnitService.
func (service *unitServiceImpl) FindByID(ctx context.Context, id string) response.ServiceResponse[*dto.UnitDTO] {
	repositoryResponse := service.UnitRepository.FindByID(ctx, id)
	if repositoryResponse.Error != nil {
		if errors.Is(repositoryResponse.Error, gorm.ErrRecordNotFound) {
			return response.ServiceNotFound(response.ServiceResponseOptions[*dto.UnitDTO]{
				Error:   repositoryResponse.Error,
				Message: repositoryResponse.Message,
			})
		}
		return response.ServiceInternalServerError(response.ServiceResponseOptions[*dto.UnitDTO]{
			Error:   repositoryResponse.Error,
			Message: repositoryResponse.Message,
		})
	}
	return response.ServiceOK(response.ServiceResponseOptions[*dto.UnitDTO]{
		Message: "successfully get unit",
		Data:    dto.ToUnit(repositoryResponse.Data),
	})
}
