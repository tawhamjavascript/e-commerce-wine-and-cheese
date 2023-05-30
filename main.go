package main

import (
	"e-commerce/config"
	"e-commerce/db"
	"e-commerce/router"
	"e-commerce/service/messagesHttp"
	"github.com/gofiber/fiber/v2"
)

func main() {
	configuration := config.LoadConfig()
	db.Connect(configuration)
	messagesHttp.InitializerMessage()
	app := fiber.New()
	router.Application(app)
	app.Listen(":8000")

}
