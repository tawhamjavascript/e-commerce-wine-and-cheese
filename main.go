package main

import (
	"e-commerce/config"
	"e-commerce/db"
	"e-commerce/router"
	"e-commerce/service/messagesHttp"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	//Load the configuration file
	configuration := config.LoadConfig()

	// This code connects to the database.
	// The configuration object contains information about the database.
	db.Connect(configuration)

	messagesHttp.InitializerMessage()


	// app is a reference to the main Fiber application instance.
	app := fiber.New()

	// Create a new instance of a cors middleware.
	app.Use(cors.New())

	// The app.Use method adds middleware to the application stack.
	router.Application(app)
	
	app.Listen(":8000")

}
