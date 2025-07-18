package infra

import (
	"fmt"
	"log"
	"sync"

	"github.com/bagusyanuar/go-erp/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func InitDB(cfg *config.DBConfig) *gorm.DB {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName,
		)
		var err error
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		log.Println("✅ Connected to PostgreSQL database successfully")
	})
	return DB
}
