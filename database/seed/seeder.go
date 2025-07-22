package seed

import "gorm.io/gorm"

func Seed(db *gorm.DB) {
	RoleSeeder(db)
	PermissionSeeder(db)
	RolePermissionSeeder(db)
	UserSeeder(db)
	FeatureSeeder(db)
}
