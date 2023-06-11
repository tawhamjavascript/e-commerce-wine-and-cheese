package clientService

import (
	"e-commerce/model"
	"e-commerce/service/messagesHttp"

	clientRepository "e-commerce/repository/client"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



func GetAllProductCheckout(c *fiber.Ctx) (*[]model.ProductView, *messagesHttp.MessageErro){
	id, _ := primitive.ObjectIDFromHex(c.Locals("id").(string))
	products, err := clientRepository.GetAllProductCheckout(c.Context(), &id)
	if err != nil {
		return nil, messagesHttp.GetError(err)
	}
	return products, nil
	
}
