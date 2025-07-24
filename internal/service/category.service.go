package service

import (
	"context"

	"github.com/bagusyanuar/go-erp/internal/delivery/request"
	"github.com/bagusyanuar/go-erp/internal/domain/dto"
	"github.com/bagusyanuar/go-erp/internal/domain/repository"
	"github.com/bagusyanuar/go-erp/pkg/lib/response"
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
		Message: "successfully get units",
		Data:    &data,
		Meta:    repositoryResponse.Meta,
	})
}

// FindByID implements CategoryService.
func (service *categoryServiceImpl) FindByID(ctx context.Context, id string) response.ServiceResponse[*dto.CategoryDTO] {
	panic("unimplemented")
}

// Create implements CategoryService.
func (service *categoryServiceImpl) Create(ctx context.Context, schema *request.CategorySchema) response.ServiceResponse[any] {
	panic("unimplemented")
}
