package main

import (
	"github.com/zulkarnen-force/fiber-starter/config"
	"github.com/zulkarnen-force/fiber-starter/router"
)

func main() {
	config.LoadConfig()

    _, err := config.InitDB()
	if err != nil {
		panic("failed to connect database")
	}

	// Initialize Fiber app
	app := config.InitFiber()



	// Setup routes
	router.Setup(app)

	// Start the server
	if err := app.Listen(":3001"); err != nil {
		panic(err)
	}
}
