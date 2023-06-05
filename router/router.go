package router

import (
	"e-commerce/controller"
	"e-commerce/middleware"

	"github.com/gofiber/fiber/v2"
)

func Application(app *fiber.App) {
	app.Post("/vendedor/cadastrar", controller.AddVendor)
	app.Post("/vendedor/login", controller.LoginVendor,  middleware.CreateToken)
	app.Post("/vendedor/produto/cadastrar", middleware.Auth, controller.AddProductVendor)
	app.Put("/vendedor/produto/editar/:idProduct", middleware.Auth, controller.UpdateVendorProduct)
	app.Delete("/vendedor/produto/deletar/:idProduct", middleware.Auth, controller.DeleteVendorProduct)
}
