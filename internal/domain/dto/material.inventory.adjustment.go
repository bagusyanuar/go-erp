package dto

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/bagusyanuar/go-erp/internal/domain/entity"
)

type DateOnly time.Time

func (d DateOnly) MarshalJSON() ([]byte, error) {
	formatted := time.Time(d).Format("2006-01-02")
	return json.Marshal(formatted)
}

func (d *DateOnly) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*d = DateOnly(t)
	return nil
}

func (d DateOnly) ToTime() time.Time {
	return time.Time(d)
}

type MaterialInventoryAdjustmentDTO struct {
	ID        string           `json:"id"`
	Material  *BaseMaterialDTO `json:"material"`
	Unit      *UnitDTO         `json:"unit"`
	Date      DateOnly         `json:"date"`
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
		Date:      DateOnly(e.Date),
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
