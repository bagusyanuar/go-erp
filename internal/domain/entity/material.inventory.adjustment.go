package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type MaterialInventoryAdjustment struct {
	ID         uuid.UUID
	MaterialID *uuid.UUID
	UnitID     *uuid.UUID
	Date       time.Time `gorm:"type:date"`
	Type       string
	Quantity   decimal.Decimal `gorm:"type:numeric(15,2);default:0"`
	AuthorID   *uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
	Material   *Material `gorm:"foreignKey:MaterialID"`
	Unit       *Unit     `gorm:"foreignKey:UnitID"`
	Author     *User     `gorm:"foreignKey:AuthorID"`
}

func (u *MaterialInventoryAdjustment) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}

func (u *MaterialInventoryAdjustment) TableName() string {
	return "material_inventories"
}
