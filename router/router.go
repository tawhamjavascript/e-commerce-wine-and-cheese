package router

import (
	"e-commerce/controller"
	"e-commerce/middleware"

	"github.com/gofiber/fiber/v2"
)

func Application(app *fiber.App) {
	app.Post("/vendedor/cadastrar", controller.AddVendor)
	app.Post("/vendedor/login", controller.LoginVendor,  middleware.CreateToken)
	path := app.Group("/")

	apiVendor:= path.Group("/vendedor", middleware.Auth)
	apiVendor.Post("/produto/cadastrar", controller.AddProductVendor)
	apiVendor.Put("/produto/editar/:idProduct", controller.UpdateVendorProduct)
	apiVendor.Delete("/produto/deletar/:idProduct",controller.DeleteVendorProduct)
	apiVendor.Get("/produto/listar", controller.GetAllVendorProduct)
	apiVendor.Get("/produto/:id", controller.GetVendorProduct)
}
