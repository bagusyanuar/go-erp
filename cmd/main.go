package main

import (
	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/internal/http"
	"github.com/bagusyanuar/go-erp/internal/infra"
)

func main() {
	viper := config.NewViper()
	logger := infra.InitLogger()
	defer logger.Sync()

	config.NewJWTManager(viper)
	http.Start()
}
