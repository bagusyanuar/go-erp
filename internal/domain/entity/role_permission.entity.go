package entity

import "github.com/google/uuid"

type RolePermission struct {
	RoleID       uuid.UUID `gorm:"primaryKey;"`
	PermissionID uuid.UUID `gorm:"primaryKey;"`
}

func (RolePermission) TableName() string {
	return "role_permissions"
}
