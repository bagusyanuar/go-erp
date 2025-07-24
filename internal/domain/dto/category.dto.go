package dto

import "github.com/bagusyanuar/go-erp/internal/domain/entity"

type CategoryDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ToCategory(entity *entity.Category) *CategoryDTO {
	return &CategoryDTO{
		ID:   entity.ID.String(),
		Name: entity.Name,
	}
}

func ToCategories(entities []entity.Category) []CategoryDTO {
	categories := make([]CategoryDTO, 0)
	for _, entity := range entities {
		e := *ToCategory(&entity)
		categories = append(categories, e)
	}
	return categories
}
