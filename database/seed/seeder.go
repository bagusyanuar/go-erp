package seed

import "gorm.io/gorm"

func Seed(db *gorm.DB) {
	UserSeeder(db)
	FeatureSeeder(db)
}
