package clientService

import (
	"e-commerce/model"
	clientRepository "e-commerce/repository/client"
	"e-commerce/service/messagesHttp"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetAllProductShoppingCart(c *fiber.Ctx) (*[]model.ProductView, *messagesHttp.MessageErro) {
	idClient, _ := primitive.ObjectIDFromHex(c.Locals("id").(string))

	products, err := clientRepository.GetAllProductShoppingCart(c.Context(), &idClient)
	if err != nil {
		return nil, messagesHttp.GetError(err)
	}
	return products, nil

} 
