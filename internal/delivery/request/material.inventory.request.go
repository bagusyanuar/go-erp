package request

import "github.com/google/uuid"

type MaterialInventorySchema struct {
	MaterialID uuid.UUID `json:"material_id" validate:"required,uuid4"`
	UnitID     uuid.UUID `json:"unit_id" validate:"required,uuid4"`
}

type MaterialInventoryQuery struct {
	Param string `json:"param" query:"param"`
	QueryPagination
	QuerySort
}
