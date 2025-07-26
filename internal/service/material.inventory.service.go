package service

import (
	"context"

	"github.com/bagusyanuar/go-erp/internal/delivery/request"
	"github.com/bagusyanuar/go-erp/internal/domain/dto"
	"github.com/bagusyanuar/go-erp/internal/domain/entity"
	"github.com/bagusyanuar/go-erp/internal/domain/repository"
	"github.com/bagusyanuar/go-erp/pkg/exception"
	"github.com/bagusyanuar/go-erp/pkg/lib"
	"github.com/bagusyanuar/go-erp/pkg/lib/response"
	"github.com/shopspring/decimal"
)

type (
	MaterialInventoryService interface {
		Create(ctx context.Context, schema *request.MaterialInventorySchema) response.ServiceResponse[any]
		FindAll(ctx context.Context, queryParams *request.MaterialInventoryQuery) response.ServiceResponse[*[]dto.MaterialInventoryDTO]
	}

	materialInventoryServiceImpl struct {
		MaterialInventoryRepository repository.MaterialInventoryRepository
	}
)

func NewMaterialInventoryService(materialInventoryRepository repository.MaterialInventoryRepository) MaterialInventoryService {
	return &materialInventoryServiceImpl{
		MaterialInventoryRepository: materialInventoryRepository,
	}
}

// Create implements MaterialInventoryService.
func (service *materialInventoryServiceImpl) Create(ctx context.Context, schema *request.MaterialInventorySchema) response.ServiceResponse[any] {
	userID, ok := lib.GetUserIDSafe(ctx)
	if !ok {
		return response.ServiceUnauthorized(response.ServiceResponseOptions[any]{
			Error:   exception.ErrInvalidUserFormat,
			Message: exception.ErrInvalidUserFormat.Error(),
		})
	}
	materialID := schema.MaterialID
	unitID := schema.UnitID

	data := &entity.MaterialInventory{
		MaterialID: &materialID,
		UnitID:     &unitID,
		Quantity:   decimal.NewFromInt(0),
		ModifiedBy: &userID,
	}
	repositoryResponse := service.MaterialInventoryRepository.Create(ctx, data)
	if repositoryResponse.Error != nil {
		return response.ServiceInternalServerError(response.ServiceResponseOptions[any]{
			Error:   repositoryResponse.Error,
			Message: repositoryResponse.Message,
		})
	}
	return response.ServiceCreated(response.ServiceResponseOptions[any]{
		Message: "successfully create material inventory",
	})
}

// FindAll implements MaterialInventoryService.
func (service *materialInventoryServiceImpl) FindAll(ctx context.Context, queryParams *request.MaterialInventoryQuery) response.ServiceResponse[*[]dto.MaterialInventoryDTO] {
	repositoryResponse := service.MaterialInventoryRepository.FindAll(ctx, queryParams)
	if repositoryResponse.Error != nil {
		return response.ServiceInternalServerError(response.ServiceResponseOptions[*[]dto.MaterialInventoryDTO]{
			Error:   repositoryResponse.Error,
			Message: repositoryResponse.Message,
		})
	}
	data := dto.ToMaterialInventories(repositoryResponse.Data)
	return response.ServiceOK(response.ServiceResponseOptions[*[]dto.MaterialInventoryDTO]{
		Message: "successfully get material inventories",
		Data:    &data,
		Meta:    repositoryResponse.Meta,
	})
}
