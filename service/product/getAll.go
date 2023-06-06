package productService

import (
	"e-commerce/model"
	productRepository "e-commerce/repository/product"
	"e-commerce/service/messagesHttp"

	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) (*[]model.ProductView, *messagesHttp.MessageErro) {
	cursor, err := productRepository.GetAll(c.Context())
	if err != nil {
		return nil, messagesHttp.GetError(err)
	}
	var products []model.ProductView
	for cursor.Next(c.Context()) {
		var product model.ProductView
		cursor.Decode(&product)
		products = append(products, product)
	}
	return &products, nil
}