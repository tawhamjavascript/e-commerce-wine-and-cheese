package clientService

import (
	"e-commerce/service/messagesHttp"
	clientRepository "e-commerce/repository/client"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func RemoveProductShoppingCart(c *fiber.Ctx) *messagesHttp.MessageErro{
	idProduct, err := primitive.ObjectIDFromHex(c.Params("idProduct"))
	if err != nil {
		return messagesHttp.GetError(err)
	}
	idClient, _ := primitive.ObjectIDFromHex(c.Locals("id").(string))
	
	err = clientRepository.RemoveItemShoppingCart(c.Context(), &idProduct, &idClient)
	if err != nil {
		return messagesHttp.GetError(err)
	}
	return nil
}