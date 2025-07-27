package request

import (
	"github.com/google/uuid"
)

type MaterialInventoryAdjustmentSchema struct {
	MaterialID uuid.UUID `json:"material_id" validate:"required,uuid4"`
	UnitID     uuid.UUID `json:"unit_id" validate:"required,uuid4"`
	Quantity   int64     `json:"quantity" validate:"required,gt=0"`
	Date       string    `json:"date" validate:"required"`
	Type       string    `json:"type" validate:"required"`
}

type MaterialInventoryAdjustmetQuery struct {
	Param string `json:"param" query:"param"`
	QueryPagination
	QuerySort
}
