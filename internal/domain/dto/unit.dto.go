package dto

import "github.com/bagusyanuar/go-erp/internal/domain/entity"

type UnitDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ToUnit(entity *entity.Unit) *UnitDTO {
	return &UnitDTO{
		ID:   entity.ID.String(),
		Name: entity.Name,
	}
}

func ToUnits(entities []entity.Unit) []UnitDTO {
	units := make([]UnitDTO, 0)
	for _, entity := range entities {
		e := *ToUnit(&entity)
		units = append(units, e)
	}
	return units
}
