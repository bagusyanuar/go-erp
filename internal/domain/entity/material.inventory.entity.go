package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type MaterialInventory struct {
	ID         uuid.UUID
	MaterialID *uuid.UUID
	UnitID     *uuid.UUID
	Quantity   decimal.Decimal `gorm:"type:numeric(15,2);default:0"`
	ModifiedBy *uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
	Material   *Material `gorm:"foreignKey:MaterialID"`
	Unit       *Unit     `gorm:"foreignKey:UnitID"`
}

func (u *MaterialInventory) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}

func (u *MaterialInventory) TableName() string {
	return "material_inventories"
}
