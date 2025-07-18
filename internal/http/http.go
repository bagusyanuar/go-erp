package http

import (
	"fmt"

	"go.uber.org/zap"
)

func Start(log *zap.Logger) {
	app := NewRouter(log)
	port := ":3000"

	fmt.Println("Fiber server running on", port)
	if err := app.Listen(port); err != nil {
		panic(err)
	}

}
