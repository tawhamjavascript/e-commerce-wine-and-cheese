package router

import (
	"e-commerce/controller"
	"github.com/gofiber/fiber/v2"
)

func Application(app *fiber.App) {
	app.Post("/vendedor/cadastrar", controller.AddVendor)
	app.Post("/vendedor/login", controller.LoginVendor)
	app.Post("/vendedor/:idVendor/produto/cadastrar", controller.AddProductVendor)
	app.Put("/vendedor/produto/editar/:idProduct", controller.UpdateVendorProduct)
	app.Delete("/vendedor/:idVendor/produto/deletar/:idProduct", controller.DeleteVendorProduct)
}
