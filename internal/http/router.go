package http

import (
	"github.com/bagusyanuar/go-erp/internal/http/handler"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func NewRouter(log *zap.Logger) *fiber.App {
	app := fiber.New()

	homeHandler := handler.NewHomeHandler(log)

	app.Get("/", homeHandler.Index)
	return app
}
