package main

import (
	"github.com/bagusyanuar/go-erp/internal/bootstrap"
	"github.com/bagusyanuar/go-erp/internal/http"
)

func main() {
	// viper := config.NewViper()
	// logger := infra.InitLogger()
	// defer logger.Sync()
	// dbConfig := config.NewDatabaseConnection(viper)
	// database := infra.InitDB(dbConfig)

	// config.NewJWTManager(viper)
	app := bootstrap.Init()
	http.Start(app)
}
