package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Permission) BeforeCreate(tx *gorm.DB) (err error) {

	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return
}

func (p *Permission) TableName() string {
	return "permissions"
}
