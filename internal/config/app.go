package config

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppConfig struct {
	Logger *zap.Logger
	DB     *gorm.DB
}
