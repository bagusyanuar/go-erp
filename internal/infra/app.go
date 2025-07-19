package infra

import (
	"log"

	"github.com/bagusyanuar/go-erp/internal/config"
)

func Load() *config.AppConfig {
	//load viper configuration
	viper := config.NewViper()

	//load zap logger configuration
	logger := InitLogger()
	defer logger.Sync()

	//load database configuration & start connection
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
