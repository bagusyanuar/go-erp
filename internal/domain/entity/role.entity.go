package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Permissions []Permission `gorm:"many2many:role_permissions;constraint:OnDelete:CASCADE;"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New()
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
	return
}

func (r *Role) TableName() string {
	return "roles"
}
