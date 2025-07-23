package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MaterialCategory struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (u *MaterialCategory) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}

func (u *MaterialCategory) TableName() string {
	return "material_categories"
}
