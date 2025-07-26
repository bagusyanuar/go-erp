package dto

import "github.com/bagusyanuar/go-erp/internal/domain/entity"

type BaseMaterialDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type MaterialDTO struct {
	ID         string        `json:"id"`
	Name       string        `json:"name"`
	Categories []CategoryDTO `json:"categories"`
}

func ToMaterial(entity *entity.Material) *MaterialDTO {
	dataCategories := entity.Categories
	categories := make([]CategoryDTO, 0)
	for _, category := range dataCategories {
		c := *ToCategory(&category)
		categories = append(categories, c)
	}
	return &MaterialDTO{
		ID:         entity.ID.String(),
		Name:       entity.Name,
		Categories: categories,
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
