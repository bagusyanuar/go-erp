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
	MaterialService interface {
		FindAll(ctx context.Context, queryParams *request.MaterialQuery) response.ServiceResponse[*[]dto.MaterialDTO]
		FindByID(ctx context.Context, id string) response.ServiceResponse[*dto.MaterialDTO]
		Create(ctx context.Context, schema *request.MaterialSchema) response.ServiceResponse[any]
	}

	materialServiceImpl struct {
		MaterialRepository repository.MaterialRepository
	}
)

func NewMaterialService(materialRepository repository.MaterialRepository) MaterialService {
	return &materialServiceImpl{
		MaterialRepository: materialRepository,
	}
}

// Create implements MaterialService.
func (service *materialServiceImpl) Create(ctx context.Context, schema *request.MaterialSchema) response.ServiceResponse[any] {
	name := schema.Name
	categories := schema.Categories

	entity := &entity.Material{
		Name: name,
	}
	repositoryResponse := service.MaterialRepository.Create(ctx, entity, categories)
	if repositoryResponse.Error != nil {
		return response.ServiceInternalServerError(response.ServiceResponseOptions[any]{
			Error:   repositoryResponse.Error,
			Message: repositoryResponse.Message,
		})
	}
	return response.ServiceCreated(response.ServiceResponseOptions[any]{
		Message: "successfully create material",
	})
}

// FindAll implements MaterialService.
func (service *materialServiceImpl) FindAll(ctx context.Context, queryParams *request.MaterialQuery) response.ServiceResponse[*[]dto.MaterialDTO] {
	repositoryResponse := service.MaterialRepository.FindAll(ctx, queryParams)
	if repositoryResponse.Error != nil {
		return response.ServiceInternalServerError(response.ServiceResponseOptions[*[]dto.MaterialDTO]{
			Error:   repositoryResponse.Error,
			Message: repositoryResponse.Message,
		})
	}
	data := dto.ToMaterials(repositoryResponse.Data)
	return response.ServiceOK(response.ServiceResponseOptions[*[]dto.MaterialDTO]{
		Message: "successfully get materials",
		Data:    &data,
		Meta:    repositoryResponse.Meta,
	})
}

// FindByID implements MaterialService.
func (service *materialServiceImpl) FindByID(ctx context.Context, id string) response.ServiceResponse[*dto.MaterialDTO] {
	repositoryResponse := service.MaterialRepository.FindByID(ctx, id)
	if repositoryResponse.Error != nil {
		if errors.Is(repositoryResponse.Error, gorm.ErrRecordNotFound) {
			return response.ServiceNotFound(response.ServiceResponseOptions[*dto.MaterialDTO]{
				Error:   repositoryResponse.Error,
				Message: repositoryResponse.Message,
			})
		}
		return response.ServiceInternalServerError(response.ServiceResponseOptions[*dto.MaterialDTO]{
			Error:   repositoryResponse.Error,
			Message: repositoryResponse.Message,
		})
	}
	return response.ServiceOK(response.ServiceResponseOptions[*dto.MaterialDTO]{
		Message: "successfully get material",
		Data:    dto.ToMaterial(repositoryResponse.Data),
	})
}
