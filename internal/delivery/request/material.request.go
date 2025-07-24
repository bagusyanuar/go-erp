package request

import "github.com/google/uuid"

type MaterialSchema struct {
	MaterialCategoryID uuid.UUID `json:"material_category_id" validate:"required,uuid4"`
	Name               string    `json:"name" validate:"required"`
}

type MaterialQuery struct {
	Param string `json:"param" query:"param"`
	QueryPagination
}
