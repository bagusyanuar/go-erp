package app

import (
	"fmt"

	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/internal/container"
	"github.com/bagusyanuar/go-erp/internal/http"
)

type AppContainer struct {
	Repository *container.RepositoryContainer
	Service    *container.ServiceContainer
	Handler    *container.HandlerContainer
}

func CreateContainer(cfg *config.AppConfig) *AppContainer {
	repo := container.InitRepository(cfg)
	service := container.InitService(cfg, repo)
	handler := container.InitHandler(cfg, service)

	return &AppContainer{
		Repository: repo,
		Service:    service,
		Handler:    handler,
	}
}

func Start(cfg *config.AppConfig, container *AppContainer) {
	server := http.NewRouter(cfg, container.Handler)
	port := ":3000"

	fmt.Println("Fiber server running on", port)
	if err := server.Listen(port); err != nil {
		panic(err)
	}
}
