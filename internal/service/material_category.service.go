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
	MaterialCategoryService interface {
		FindAll(ctx context.Context, queryParams *request.MaterialCategoryQuery) response.ServiceResponse[*[]dto.MaterialCategoryDTO]
		FindByID(ctx context.Context, id string) response.ServiceResponse[*dto.MaterialCategoryDTO]
		Create(ctx context.Context, schema *request.MaterialCategorySchema) response.ServiceResponse[any]
	}

	materialCategoryServiceImpl struct {
		MaterialCategoryRepository repository.MaterialCategoryRepository
	}
)

func NewMaterialCategoryService(materialCategoryRepository repository.MaterialCategoryRepository) MaterialCategoryService {
	return &materialCategoryServiceImpl{
		MaterialCategoryRepository: materialCategoryRepository,
	}
}

// Create implements MaterialCategoryService.
func (service *materialCategoryServiceImpl) Create(ctx context.Context, schema *request.MaterialCategorySchema) response.ServiceResponse[any] {
	name := schema.Name

	data := &entity.MaterialCategory{
		Name: name,
	}
	repositoryResponse := service.MaterialCategoryRepository.Create(ctx, data)
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

// FindAll implements MaterialCategoryService.
func (service *materialCategoryServiceImpl) FindAll(ctx context.Context, queryParams *request.MaterialCategoryQuery) response.ServiceResponse[*[]dto.MaterialCategoryDTO] {
	repositoryResponse := service.MaterialCategoryRepository.FindAll(ctx, queryParams)
	if repositoryResponse.Error != nil {
		return response.ServiceInternalServerError(response.ServiceResponseOptions[*[]dto.MaterialCategoryDTO]{
			Error:   repositoryResponse.Error,
			Message: repositoryResponse.Message,
		})
	}
	data := dto.ToMaterialCategories(repositoryResponse.Data)
	return response.ServiceOK(response.ServiceResponseOptions[*[]dto.MaterialCategoryDTO]{
		Message: "successfully get material categories",
		Data:    &data,
		Meta:    repositoryResponse.Meta,
	})
}

// FindByID implements MaterialCategoryService.
func (service *materialCategoryServiceImpl) FindByID(ctx context.Context, id string) response.ServiceResponse[*dto.MaterialCategoryDTO] {
	repositoryResponse := service.MaterialCategoryRepository.FindByID(ctx, id)
	if repositoryResponse.Error != nil {
		if errors.Is(repositoryResponse.Error, gorm.ErrRecordNotFound) {
			return response.ServiceNotFound(response.ServiceResponseOptions[*dto.MaterialCategoryDTO]{
				Error:   repositoryResponse.Error,
				Message: repositoryResponse.Message,
			})
		}
		return response.ServiceInternalServerError(response.ServiceResponseOptions[*dto.MaterialCategoryDTO]{
			Error:   repositoryResponse.Error,
			Message: repositoryResponse.Message,
		})
	}
	return response.ServiceOK(response.ServiceResponseOptions[*dto.MaterialCategoryDTO]{
		Message: "successfully get material category",
		Data:    dto.ToMaterialCategory(repositoryResponse.Data),
	})
}
