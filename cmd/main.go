package main

import (
	"github.com/bagusyanuar/go-erp/internal/app"
)

func main() {
	appConfig := app.Load()
	container := app.CreateContainer(appConfig)
	app.Start(appConfig, container)
}
