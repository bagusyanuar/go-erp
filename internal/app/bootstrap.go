package app

import (
	"log"

	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/internal/infra"
)

func Load() *config.AppConfig {
	viper := config.NewViper()

	logger := infra.InitLogger()
	defer logger.Sync()

	dbConfig := config.NewDatabaseConnection(viper)
	db := infra.InitDB(dbConfig)

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
