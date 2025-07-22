package seed

import (
	"errors"
	"log"

	"github.com/bagusyanuar/go-erp/internal/domain/entity"
	"gorm.io/gorm"
)

func RoleSeeder(db *gorm.DB) {
	if !db.Migrator().HasTable("roles") {
		log.Println("⛔ Table roles not found, seeding cancelled.")
		return
	}

	roles := []string{
		"superdev",
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		for _, r := range roles {
			var role entity.Role
			err := tx.Where("name = ?", r).First(&role).Error

			data := entity.Role{
				Name: r,
			}
			switch {
			case errors.Is(err, gorm.ErrRecordNotFound):
				if err := tx.Create(&data).Error; err != nil {
					log.Printf("failed to insert role: %s", r)
					return err
				}
				log.Printf("inserted role: %s", r)
			case err != nil:
				log.Printf("error checking role %s: %v", r, err)
				return err
			default:
				if err := tx.Model(&role).Updates(data).Error; err != nil {
					log.Printf("failed to update role: %s", r)
					return err
				}
				log.Printf("updated role: %s", r)
			}
		}
		return nil
	})

	if err != nil {
		log.Printf("failed seed reoles: %v", err)
	} else {
		log.Println("successfully seed roles")
	}
}
