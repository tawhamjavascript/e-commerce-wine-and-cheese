package router

import (
	"e-commerce/controller"
	"e-commerce/middleware"

	"github.com/gofiber/fiber/v2"
)

func Application(app *fiber.App) {
	app.Post("/vendedor/cadastrar", controller.AddVendor)
	app.Post("/cliente/cadastrar", controller.RegisterClient)
	app.Post("/cliente/login", controller.LoginClient, middleware.CreateToken)
	app.Post("/vendedor/login", controller.LoginVendor,  middleware.CreateToken)
	
	path := app.Group("/")

	//-------------------------------------------------------------------------------------
	apiVendor := path.Group("/vendedor", middleware.Auth)
	apiVendor.Get("/produto/vendidos", controller.GetAllVendorProductSold)
	apiVendor.Post("/produto/cadastrar", controller.AddProductVendor)
	apiVendor.Put("/produto/editar/:idProduct", controller.UpdateVendorProduct)
	apiVendor.Delete("/produto/deletar/:idProduct",controller.DeleteVendorProduct)
	apiVendor.Get("/produto", controller.GetAllVendorProduct)
	apiVendor.Get("/produto/:id", controller.GetVendorProduct)
	
	//-------------------------------------------------------------------------------------

	apiClient := path.Group("/cliente", middleware.Auth)
	apiClient.Get("/produto", controller.GetAllProduct)
	apiClient.Post("/carrinho", controller.AddShoppingCart)
	apiClient.Delete("/carrinho/:idProduct", controller.RemoveItemShoppingCart)
	apiClient.Get("/carrinho", controller.GetAllProductShoppingCart)
	apiClient.Post("/carrinho/comprar", controller.Checkout)
	apiClient.Get("/comprados", controller.GetAllProductCheckout)


}
