package service

import (
	"context"
	"time"

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
	MaterialInventoryAdjustmentService interface {
		Create(ctx context.Context, schema *request.MaterialInventoryAdjustmentSchema) response.ServiceResponse[any]
		FindAll(ctx context.Context, queryParams *request.MaterialInventoryAdjustmetQuery) response.ServiceResponse[*[]dto.MaterialInventoryAdjustmentDTO]
	}

	materialInventoryAdjustmentServiceImpl struct {
		MaterialInventoryAdjustmentRepository repository.MaterialInventoryAdjustmentRepository
	}
)

func NewMaterialInventoryAdjustmentService(materialInventoryAdjustmentRepository repository.MaterialInventoryAdjustmentRepository) MaterialInventoryAdjustmentService {
	return &materialInventoryAdjustmentServiceImpl{
		MaterialInventoryAdjustmentRepository: materialInventoryAdjustmentRepository,
	}
}

// Create implements MaterialInventoryAdjustmentService.
func (service *materialInventoryAdjustmentServiceImpl) Create(ctx context.Context, schema *request.MaterialInventoryAdjustmentSchema) response.ServiceResponse[any] {
	userID, ok := lib.GetUserIDSafe(ctx)
	if !ok {
		return response.ServiceUnauthorized(response.ServiceResponseOptions[any]{
			Error:   exception.ErrInvalidUserFormat,
			Message: exception.ErrInvalidUserFormat.Error(),
		})
	}

	materialID := schema.MaterialID
	unitID := schema.UnitID

	dateVal, err := time.Parse("2006-01-02", schema.Date)
	if err != nil {
		return response.ServiceBadRequest(response.ServiceResponseOptions[any]{
			Error:   exception.ErrBadRequest,
			Message: err.Error(),
		})
	}

	data := &entity.MaterialInventoryAdjustment{
		MaterialID: &materialID,
		UnitID:     &unitID,
		Date:       dateVal,
		Quantity:   decimal.NewFromInt(schema.Quantity),
		AuthorID:   &userID,
		Type:       schema.Type,
	}

	repositoryResponse := service.MaterialInventoryAdjustmentRepository.Create(ctx, data)
	if repositoryResponse.Error != nil {
		return response.ServiceInternalServerError(response.ServiceResponseOptions[any]{
			Error:   repositoryResponse.Error,
			Message: repositoryResponse.Message,
		})
	}
	return response.ServiceCreated(response.ServiceResponseOptions[any]{
		Message: "successfully create material inventory adjustment",
	})
}

// FindAll implements MaterialInventoryAdjustmentService.
func (service *materialInventoryAdjustmentServiceImpl) FindAll(ctx context.Context, queryParams *request.MaterialInventoryAdjustmetQuery) response.ServiceResponse[*[]dto.MaterialInventoryAdjustmentDTO] {
	repositoryResponse := service.MaterialInventoryAdjustmentRepository.FindAll(ctx, queryParams)
	if repositoryResponse.Error != nil {
		return response.ServiceInternalServerError(response.ServiceResponseOptions[*[]dto.MaterialInventoryAdjustmentDTO]{
			Error:   repositoryResponse.Error,
			Message: repositoryResponse.Message,
		})
	}
	data := dto.ToMaterialInventoryAdjustments(repositoryResponse.Data)
	return response.ServiceOK(response.ServiceResponseOptions[*[]dto.MaterialInventoryAdjustmentDTO]{
		Message: "successfully get material inventory adjustments",
		Data:    &data,
		Meta:    repositoryResponse.Meta,
	})
}
