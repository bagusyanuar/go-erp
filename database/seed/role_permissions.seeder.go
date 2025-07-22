package seed

import (
	"fmt"
	"log"

	"github.com/bagusyanuar/go-erp/internal/domain/entity"
	"gorm.io/gorm"
)

func RolePermissionSeeder(db *gorm.DB) {
	if !db.Migrator().HasTable("role_permissions") {
		log.Println("⛔ Table role_permissions not found, seeding cancelled.")
		return
	}

	permissions := []string{}
	for _, feature := range Features {
		for _, action := range Actions {
			permission := fmt.Sprintf("%s.%s", feature, action)
			permissions = append(permissions, permission)
		}
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		var role entity.Role
		if err := tx.First(&role, "name = ?", "superdev").Error; err != nil {
			log.Printf("failed to find role")
			return err
		}

		var dataPermissions []entity.Permission
		if err := tx.Where("name IN ?", permissions).Find(&dataPermissions).Error; err != nil {
			log.Printf("failed to query permissions")
			return err
		}

		if err := tx.Model(&role).Association("Permissions").Replace(dataPermissions); err != nil {
			log.Printf("failed to create role_permissions seeder")
			return err
		}

		return nil
	})

	if err != nil {
		log.Printf("failed seed role permissions: %v", err)
	} else {
		log.Println("successfully seed role permissions")
	}

}
