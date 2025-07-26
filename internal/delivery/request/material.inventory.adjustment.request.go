package request

import "github.com/google/uuid"

type MaterialInventoryAdjustmentSchema struct {
	MaterialID uuid.UUID `json:"material_id" validate:"required,uuid4"`
	UnitID     uuid.UUID `json:"unit_id" validate:"required,uuid4"`
	Quantity   int       `json:"quantity" validate:"required,gt=0"`
}

type MaterialInventoryAdjustmetQuery struct {
	Param string `json:"param" query:"param"`
	QueryPagination
	QuerySort
}
