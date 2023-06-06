package productService

import (
	"e-commerce/model"
	productRepositoy "e-commerce/repository/product"
	"e-commerce/service/messagesHttp"
	"e-commerce/validate"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateProduct(c *fiber.Ctx) *messagesHttp.MessageErro{
	var productUpdate model.RegisterUpdateProduct
	idProduct, err := primitive.ObjectIDFromHex(c.Params("idProduct"))
	if err != nil {
		return messagesHttp.GetError(err)
	}
	if err := c.BodyParser(&productUpdate); err != nil {
		return messagesHttp.GetError(err)
	}
	
	err = validate.ValidateUpdateProduct(&productUpdate)
	if err != nil {
		return messagesHttp.GetError(err)
	}
	err = productRepositoy.UpdateProduct(c.Context(), &idProduct, &productUpdate)
	if err != nil {
		return messagesHttp.GetError(err)
	}
	return nil

}