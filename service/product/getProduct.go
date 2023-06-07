package productService

import (
	"e-commerce/model"
	productRepository "e-commerce/repository/product"
	"e-commerce/service/messagesHttp"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetProduct(c *fiber.Ctx) (*model.ProductView, *messagesHttp.MessageErro){
	id_product := c.Params("id")
	id_product_object, err := primitive.ObjectIDFromHex(id_product)
	if err != nil {
		return nil, messagesHttp.GetError(err)
	}
	result := productRepository.GetProduct(c.Context(), &id_product_object)
	if result.Err() != nil {
		return nil, messagesHttp.GetError(result.Err())
	}
	var product model.ProductView
	result.Decode(&product)
	return &product, nil
	
}