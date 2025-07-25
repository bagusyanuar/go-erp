package service

import (
	"context"
	"fmt"
	"log"

	"github.com/bagusyanuar/go-erp/internal/delivery/request"
	"github.com/bagusyanuar/go-erp/internal/domain/entity"
	"github.com/bagusyanuar/go-erp/internal/domain/repository"
	"github.com/bagusyanuar/go-erp/pkg/constant"
	"github.com/bagusyanuar/go-erp/pkg/lib/response"
	"github.com/shopspring/decimal"
)

type (
	MaterialInventoryService interface {
		Create(ctx context.Context, schema *request.MaterialInventorySchema) response.ServiceResponse[any]
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
	userID := ctx.Value(constant.UserIDKey)
	if userIDStr, ok := userID.(string); ok {
		fmt.Println("user ID:", userIDStr)
		// pakai userIDStr di sini
	} else {
		// userID tidak ditemukan, error handling
	}
	log.Println(userID)
	materialID := schema.MaterialID
	unitID := schema.UnitID

	data := &entity.MaterialInventory{
		MaterialID: &materialID,
		UnitID:     &unitID,
		Quantity:   decimal.NewFromInt(0),
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
