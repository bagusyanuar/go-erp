package seed

import (
	"errors"
	"fmt"
	"log"

	"github.com/bagusyanuar/go-erp/internal/domain/entity"
	"gorm.io/gorm"
)

func PermissionSeeder(db *gorm.DB) {
	if !db.Migrator().HasTable("permissions") {
		log.Println("⛔ Table permissions not found, seeding cancelled.")
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
		for _, p := range permissions {
			var permission entity.Permission
			err := tx.Where("name = ?", p).First(&permission).Error

			data := entity.Permission{
				Name: p,
			}
			switch {
			case errors.Is(err, gorm.ErrRecordNotFound):
				if err := tx.Create(&data).Error; err != nil {
					log.Printf("failed to insert permission: %s", p)
					return err
				}
				log.Printf("inserted permission: %s", p)
			case err != nil:
				log.Printf("error checking permission %s: %v", p, err)
				return err
			default:
				if err := tx.Model(&permission).Updates(data).Error; err != nil {
					log.Printf("failed to update permission: %s", p)
					return err
				}
				log.Printf("updated permission: %s", p)
			}
		}
		return nil
	})

	if err != nil {
		log.Printf("failed seed permissions: %v", err)
	} else {
		log.Println("successfully seed permissions")
	}

}
