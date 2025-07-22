package app

import (
	"fmt"

	"github.com/bagusyanuar/go-erp/internal/app/di"
	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/internal/http"
)

type AppContainer struct {
	Repository *di.RepositoryContainer
	Service    *di.ServiceContainer
	Handler    *di.HandlerContainer
}

func CreateContainer(cfg *config.AppConfig) *AppContainer {
	repo := di.InitRepository(cfg)
	service := di.InitService(cfg, repo)
	handler := di.InitHandler(cfg, service)

	return &AppContainer{
		Repository: repo,
		Service:    service,
		Handler:    handler,
	}
}

func Start(cfg *config.AppConfig, container *AppContainer) {
	http.NewRouter(cfg, container.Handler)
	envPort := cfg.Viper.GetString("APP_PORT")
	port := fmt.Sprintf(":%s", envPort)

	server := cfg.App

	fmt.Println("Fiber server running on", port)
	if err := server.Listen(port); err != nil {
		panic(err)
	}
}
