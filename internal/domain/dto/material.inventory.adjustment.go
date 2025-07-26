package dto

import (
	"time"

	"github.com/bagusyanuar/go-erp/internal/domain/entity"
)

type MaterialInventoryAdjustmentDTO struct {
	ID        string           `json:"id"`
	Material  *BaseMaterialDTO `json:"material"`
	Unit      *UnitDTO         `json:"unit"`
	Date      time.Time        `json:"date"`
	Quantity  float64          `json:"quantity"`
	Author    *UserDTO         `json:"author"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

func ToMaterialInventoryAdjustment(e *entity.MaterialInventoryAdjustment) *MaterialInventoryAdjustmentDTO {
	var material *BaseMaterialDTO
	var unit *UnitDTO
	var author *UserDTO
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

	if e.Author != nil {
		author = &UserDTO{
			ID:       e.Author.ID.String(),
			Email:    e.Author.Email,
			Username: e.Author.Username,
		}
	}
	return &MaterialInventoryAdjustmentDTO{
		ID:        e.ID.String(),
		Material:  material,
		Unit:      unit,
		Quantity:  e.Quantity.Round(2).InexactFloat64(),
		Author:    author,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func ToMaterialInventoryAdjustments(entities []entity.MaterialInventoryAdjustment) []MaterialInventoryAdjustmentDTO {
	materialInventoryAdjustment := make([]MaterialInventoryAdjustmentDTO, 0)
	for _, entity := range entities {
		e := *ToMaterialInventoryAdjustment(&entity)
		materialInventoryAdjustment = append(materialInventoryAdjustment, e)
	}
	return materialInventoryAdjustment
}
