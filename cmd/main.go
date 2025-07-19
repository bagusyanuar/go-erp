package main

import (
	"github.com/bagusyanuar/go-erp/internal/http"
	"github.com/bagusyanuar/go-erp/internal/infra"
)

func main() {
	app := infra.Load()
	http.Start(app)
}
