package productService

import (
	"e-commerce/model"
	"e-commerce/service/messagesHttp"
	"log"

	productRepository "e-commerce/repository/product"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



func GetAllProductSoldVendor(c *fiber.Ctx) (*[]*model.ProductView, *messagesHttp.MessageErro) {
	id, _ := primitive.ObjectIDFromHex(c.Locals("id").(string))
	log.Println("Entrando aqui")
	products, err := productRepository.GetProductSoldVendor(c.Context(), &id)

	if err != nil {
		log.Println(err)
		return nil, messagesHttp.GetError(err)
	}
	var productsView []*model.ProductView
	for _, product := range *products {
		productView := model.ProductView{
			ID:          product.ID,
			Name:        *product.Name,
			Chapter:     *product.Chapter,
			Price:       *product.Price,
			Description: *product.Description,
			Author:      *product.Author,
			Genre:       *product.Genre,
			Publication: *product.Publication,
			Photo:       *product.Photo,
			Vendor:      *product.Vendor,
		}
		productsView = append(productsView, &productView)
	}
	log.Println("Entrando aqui")
	return &productsView, nil
}