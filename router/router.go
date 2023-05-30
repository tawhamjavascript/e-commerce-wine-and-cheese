package router

import (
	"e-commerce/controller"
	"github.com/gofiber/fiber/v2"
)

func Application(app *fiber.App) {
	app.Post("/vendedor/cadastrar", controller.AddVendor)
	app.Post("/vendedor/login", controller.LoginVendor)
}
