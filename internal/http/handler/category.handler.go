package handler

import (
	"github.com/bagusyanuar/go-erp/internal/config"
	"github.com/bagusyanuar/go-erp/internal/service"
)

type CategoryHandler struct {
	CategoryService service.CategoryService
	Config          *config.AppConfig
}

func NewCategoryHandler(categoryService service.CategoryService, cfg *config.AppConfig) *CategoryHandler {
	return &CategoryHandler{
		CategoryService: categoryService,
		Config:          cfg,
	}
}
