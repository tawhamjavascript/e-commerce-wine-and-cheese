package clientService

import (
	"e-commerce/model"
	clientRepository "e-commerce/repository/client"
	"e-commerce/service/messagesHttp"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func AddShoppingCart(c *fiber.Ctx)(*model.ProductView, *messagesHttp.MessageErro) {
	var productView model.ProductView

	err := c.BodyParser(&productView)
	if err != nil {
		return nil, messagesHttp.GetError(err)
	}
	idClient, _ := primitive.ObjectIDFromHex(c.Locals("id").(string))
	erro := clientRepository.AddShoppingCart(c.Context(), &productView, &idClient)
	if erro != nil {
		return nil, messagesHttp.GetError(erro)
	}
	return &productView, nil

}