package dto

import (
	"github.com/bagusyanuar/go-erp/internal/domain/entity"
	"github.com/shopspring/decimal"
)

type MaterialInventoryDTO struct {
	ID       string          `json:"id"`
	Material *MaterialDTO    `json:"material"`
	Unit     *UnitDTO        `json:"unit"`
	Quantity decimal.Decimal `json:"quantity"`
}

func ToMaterialInventory(entity *entity.MaterialInventory) *MaterialInventoryDTO {
	material := ToMaterial(entity.Material)
	unit := ToUnit(entity.Unit)
	return &MaterialInventoryDTO{
		ID:       entity.ID.String(),
		Material: material,
		Unit:     unit,
		Quantity: entity.Quantity,
	}
}

func ToMaterialInventories(entities []entity.MaterialInventory) []MaterialInventoryDTO {
	materialInventories := make([]MaterialInventoryDTO, 0)
	for _, entity := range entities {
		e := *ToMaterialInventory(&entity)
		materialInventories = append(materialInventories, e)
	}
	return materialInventories
}
