package dto

import (
	"time"

	"github.com/bagusyanuar/go-erp/internal/domain/entity"
)

type MaterialInventoryDTO struct {
	ID         string           `json:"id"`
	Material   *BaseMaterialDTO `json:"material"`
	Unit       *UnitDTO         `json:"unit"`
	Quantity   float64          `json:"quantity"`
	ModifiedBy *UserDTO         `json:"modified_by"`
	CreatedAt  time.Time        `json:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at"`
}

func ToMaterialInventory(e *entity.MaterialInventory) *MaterialInventoryDTO {
	var material *BaseMaterialDTO
	var unit *UnitDTO
	var modifiedBy *UserDTO
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

	if e.Modificator != nil {
		modifiedBy = &UserDTO{
			ID:       e.Modificator.ID.String(),
			Email:    e.Modificator.Email,
			Username: e.Modificator.Username,
		}
	}
	return &MaterialInventoryDTO{
		ID:         e.ID.String(),
		Material:   material,
		Unit:       unit,
		Quantity:   e.Quantity.Round(2).InexactFloat64(),
		ModifiedBy: modifiedBy,
		CreatedAt:  e.CreatedAt,
		UpdatedAt:  e.UpdatedAt,
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
