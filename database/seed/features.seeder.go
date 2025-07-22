package seed

import (
	"errors"
	"log"

	"github.com/bagusyanuar/go-erp/internal/domain/entity"
	"gorm.io/gorm"
)

func FeatureSeeder(db *gorm.DB) {
	if !db.Migrator().HasTable("features") {
		log.Println("⛔ Table features not found, seeding cancelled.")
		return
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		for _, f := range Features {
			var feature entity.Feature
			err := tx.Where("name = ?", f).First(&feature).Error

			data := entity.Feature{
				Name: f,
			}
			switch {
			case errors.Is(err, gorm.ErrRecordNotFound):
				if err := tx.Create(&data).Error; err != nil {
					log.Printf("failed to insert feature: %s", f)
					return err
				}
				log.Printf("inserted feature: %s", f)
			case err != nil:
				log.Printf("error checking feature %s: %v", f, err)
				return err
			default:
				if err := tx.Model(&feature).Updates(data).Error; err != nil {
					log.Printf("failed to update feature: %s", f)
					return err
				}
				log.Printf("updated feature: %s", f)
			}
		}
		return nil
	})

	if err != nil {
		log.Printf("failed seed features: %v", err)
	} else {
		log.Println("successfully seed features")
	}
}
