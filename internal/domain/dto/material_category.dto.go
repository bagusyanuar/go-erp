package dto

import "github.com/bagusyanuar/go-erp/internal/domain/entity"

type MaterialCategoryDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ToMaterialCategory(entity *entity.MaterialCategory) *MaterialCategoryDTO {
	return &MaterialCategoryDTO{
		ID:   entity.ID.String(),
		Name: entity.Name,
	}
}

func ToMaterialCategories(entities []entity.MaterialCategory) []MaterialCategoryDTO {
	materialCategories := make([]MaterialCategoryDTO, 0)
	for _, entity := range entities {
		e := *ToMaterialCategory(&entity)
		materialCategories = append(materialCategories, e)
	}
	return materialCategories
}
