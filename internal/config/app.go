package config

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppConfig struct {
	Logger    *zap.Logger
	DB        *gorm.DB
	Validator *validator.Validate
	JWT       *JWTConfig
}
