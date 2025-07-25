package dto

import "github.com/bagusyanuar/go-erp/internal/domain/entity"

type MaterialDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ToMaterial(entity *entity.Material) *MaterialDTO {
	return &MaterialDTO{
		ID:   entity.ID.String(),
		Name: entity.Name,
	}
}

func ToMaterials(entities []entity.Material) []MaterialDTO {
	materials := make([]MaterialDTO, 0)
	for _, entity := range entities {
		e := *ToMaterial(&entity)
		materials = append(materials, e)
	}
	return materials
}
