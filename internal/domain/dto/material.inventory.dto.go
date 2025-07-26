package dto

import (
	"github.com/bagusyanuar/go-erp/internal/domain/entity"
)

type MaterialInventoryDTO struct {
	ID       string           `json:"id"`
	Material *BaseMaterialDTO `json:"material"`
	Unit     *UnitDTO         `json:"unit"`
	Quantity float64          `json:"quantity"`
}

func ToMaterialInventory(e *entity.MaterialInventory) *MaterialInventoryDTO {
	var material *BaseMaterialDTO
	var unit *UnitDTO
	if e.Material != nil {
		material = &BaseMaterialDTO{
			ID:   e.Material.ID.String(),
			Name: e.Material.Name,
		}
	}

	if e.Unit != nil {
		unit = &UnitDTO{
			ID:   e.Unit.ID.String(),
			Name: e.Unit.Name,
		}
	}
	return &MaterialInventoryDTO{
		ID:       e.ID.String(),
		Material: material,
		Unit:     unit,
		Quantity: e.Quantity.Round(2).InexactFloat64(),
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
