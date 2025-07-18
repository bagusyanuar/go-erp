package http

import "fmt"

func Start() {
	app := NewRouter()
	port := ":3000"

	fmt.Println("Fiber server running on", port)
	if err := app.Listen(port); err != nil {
		panic(err)
	}

}
