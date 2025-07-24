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
	CategoryService interface {
		FindAll(ctx context.Context, queryParams *request.CategoryQuery) response.ServiceResponse[*[]dto.CategoryDTO]
		FindByID(ctx context.Context, id string) response.ServiceResponse[*dto.CategoryDTO]
		Create(ctx context.Context, schema *request.CategorySchema) response.ServiceResponse[any]
	}

	categoryServiceImpl struct {
		CategoryRepository repository.CategoryRepository
	}
)

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &categoryServiceImpl{
		CategoryRepository: categoryRepository,
	}
}

// FindAll implements CategoryService.
func (service *categoryServiceImpl) FindAll(ctx context.Context, queryParams *request.CategoryQuery) response.ServiceResponse[*[]dto.CategoryDTO] {
	repositoryResponse := service.CategoryRepository.FindAll(ctx, queryParams)
	if repositoryResponse.Error != nil {
		return response.ServiceInternalServerError(response.ServiceResponseOptions[*[]dto.CategoryDTO]{
			Error:   repositoryResponse.Error,
			Message: repositoryResponse.Message,
		})
	}
	data := dto.ToCategories(repositoryResponse.Data)
	return response.ServiceOK(response.ServiceResponseOptions[*[]dto.CategoryDTO]{
		Message: "successfully get categories",
		Data:    &data,
		Meta:    repositoryResponse.Meta,
	})
}

// FindByID implements CategoryService.
func (service *categoryServiceImpl) FindByID(ctx context.Context, id string) response.ServiceResponse[*dto.CategoryDTO] {
	repositoryResponse := service.CategoryRepository.FindByID(ctx, id)
	if repositoryResponse.Error != nil {
		if errors.Is(repositoryResponse.Error, gorm.ErrRecordNotFound) {
			return response.ServiceNotFound(response.ServiceResponseOptions[*dto.CategoryDTO]{
				Error:   repositoryResponse.Error,
				Message: repositoryResponse.Message,
			})
		}
		return response.ServiceInternalServerError(response.ServiceResponseOptions[*dto.CategoryDTO]{
			Error:   repositoryResponse.Error,
			Message: repositoryResponse.Message,
		})
	}
	return response.ServiceOK(response.ServiceResponseOptions[*dto.CategoryDTO]{
		Message: "successfully get category",
		Data:    dto.ToCategory(repositoryResponse.Data),
	})
}

// Create implements CategoryService.
func (service *categoryServiceImpl) Create(ctx context.Context, schema *request.CategorySchema) response.ServiceResponse[any] {
	name := schema.Name

	data := &entity.Category{
		Name: name,
	}
	repositoryResponse := service.CategoryRepository.Create(ctx, data)
	if repositoryResponse.Error != nil {
		return response.ServiceInternalServerError(response.ServiceResponseOptions[any]{
			Error:   repositoryResponse.Error,
			Message: repositoryResponse.Message,
		})
	}
	return response.ServiceCreated(response.ServiceResponseOptions[any]{
		Message: "successfully create category",
	})
}
