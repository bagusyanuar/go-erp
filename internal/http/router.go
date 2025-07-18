package http

import (
	"github.com/bagusyanuar/go-erp/internal/http/handler"
	"github.com/gofiber/fiber/v2"
)

func NewRouter() *fiber.App {
	app := fiber.New()

	homeHandler := handler.NewHomeHandler()

	app.Get("/", homeHandler.Index)
	return app
}
