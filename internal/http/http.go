package http

import (
	"fmt"

	"github.com/bagusyanuar/go-erp/internal/bootstrap"
	"github.com/bagusyanuar/go-erp/internal/bootstrap/container"
)

func Start(cfg *bootstrap.AppConfig) {
	repo := container.InitRepository(cfg)
	svc := container.InitService(cfg, repo)
	handler := container.InitHandler(cfg, svc)
	app := NewRouter(cfg, handler)
	port := ":3000"

	fmt.Println("Fiber server running on", port)
	if err := app.Listen(port); err != nil {
		panic(err)
	}

}
