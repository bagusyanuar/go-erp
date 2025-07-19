package infra

import (
	"log"

	"github.com/bagusyanuar/go-erp/internal/config"
)

func Load() *config.AppConfig {
	viper := config.NewViper()

	logger := InitLogger()
	defer logger.Sync()

	dbConfig := config.NewDatabaseConnection(viper)
	db := InitDB(dbConfig)

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get database instance: %v", err)
	}

	defer sqlDB.Close()

	return &config.AppConfig{
		Logger: logger,
		DB:     db,
	}
}
