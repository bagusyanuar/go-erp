package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Feature struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (f *Feature) BeforeCreate(tx *gorm.DB) (err error) {
	if f.ID == uuid.Nil {
		f.ID = uuid.New()
	}
	f.CreatedAt = time.Now()
	f.UpdatedAt = time.Now()
	return
}

func (f *Feature) TableName() string {
	return "features"
}
