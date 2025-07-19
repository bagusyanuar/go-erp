package bootstrap

import (
	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/internal/infra"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppConfig struct {
	Logger *zap.Logger
	DB     *gorm.DB
}

func Init() *AppConfig {
	viper := config.NewViper()
	logger := infra.InitLogger()
	defer logger.Sync()
	dbConfig := config.NewDatabaseConnection(viper)
	database := infra.InitDB(dbConfig)

	return &AppConfig{
		Logger: logger,
		DB:     database,
	}
}
