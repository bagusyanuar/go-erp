package app

import (
	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/internal/infra"
)

func Load() *config.AppConfig {
	viper := config.NewViper()

	logger := infra.InitLogger()
	defer logger.Sync()

	dbConfig := config.NewDatabaseConnection(viper)
	db := infra.InitDB(dbConfig)

	validator := config.NewValidator()

	jwtConfig := config.NewJWTManager(viper)

	app := config.NewFiber(viper)

	return &config.AppConfig{
		App:       app,
		Logger:    logger,
		DB:        db,
		Validator: validator,
		JWT:       jwtConfig,
		Viper:     viper,
	}
}
