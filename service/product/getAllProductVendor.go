package productService

import (
	"e-commerce/model"
	productRepository "e-commerce/repository/product"
	"e-commerce/service/messagesHttp"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetAllProductVendor(c *fiber.Ctx) (*[]*model.ProductView, *messagesHttp.MessageErro) {
	idVendor, _ := primitive.ObjectIDFromHex(c.Locals("id").(string))
	products, err := productRepository.GetAllProductVendor(c.Context(), &idVendor)
	if err != nil {
		return nil, messagesHttp.GetError(err)
	}
	return products, nil

}