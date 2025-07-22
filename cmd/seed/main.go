package main

import (
	"github.com/bagusyanuar/go-erp/database/seed"
	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/internal/infra"
)

func main() {
	viper := config.NewViper()
	dbConfig := config.NewDatabaseConnection(viper)
	db := infra.InitDB(dbConfig)

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	seed.Seed(db)
}
