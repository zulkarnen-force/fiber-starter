package main

import (
	"fmt"

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
fmt.Println(":"+config.AppConfig.Port);
	// Start the server
	if err := app.Listen(":"+config.AppConfig.Port); err != nil {
		panic(err)
	} 
		fmt.Println("Server running on port ", config.AppConfig.Port)
	
	
}
