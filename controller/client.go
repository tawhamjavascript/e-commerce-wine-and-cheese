package controller

import (
	clientService "e-commerce/service/client"
	productService "e-commerce/service/product"
	"log"

	"github.com/gofiber/fiber/v2"
)



func RegisterClient(c *fiber.Ctx) error {
	message := clientService.Register(c)
	if message == nil {
		return c.Status(fiber.StatusCreated).JSON(
			fiber.Map{
				"message": "Client created",
			},
		)
	}
	return c.Status(message.Status).JSON(fiber.Map{
		"message": message.Message,
	})
}

func LoginClient(c *fiber.Ctx) error {
	log.Println("djdjdjdjjd")
	message := clientService.Login(c)
	log.Println(message, "djdjdjdjjd")
	if message != nil {
		return c.Status(message.Status).JSON(fiber.Map{
		"message": message.Message,
		})
	}
	return c.Next()	
}

func GetAllProduct(c *fiber.Ctx) error {
	products, message := productService.GetAll(c)
	if message == nil {
		return c.Status(fiber.StatusOK).JSON(products)
	}
	return c.Status(message.Status).JSON(fiber.Map{
		"message": message.Message,
	})
}

func AddShoppingCart(c *fiber.Ctx) error {
	product, message := clientService.AddShoppingCart(c)
	if message == nil {
		return c.Status(fiber.StatusOK).JSON(product)
	}
	return c.Status(message.Status).JSON(fiber.Map{
		"message": message.Message,
	})
}

func RemoveItemShoppingCart(c *fiber.Ctx) error {
	message := clientService.RemoveProductShoppingCart(c)
	if message == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Product removed",
		})
	}
	return c.Status(message.Status).JSON(fiber.Map{
		"message": message.Message,
	})
}

func GetAllProductShoppingCart(c *fiber.Ctx) error {
	products, message := clientService.GetAllProductShoppingCart(c)
	if message == nil {
		return c.Status(fiber.StatusOK).JSON(products)
	}
	return c.Status(message.Status).JSON(fiber.Map{
		"message": message.Message,
	})
}

func Checkout(c *fiber.Ctx) error {
	message := clientService.Checkout(c)
	if message == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Checkout done",
		})
	}
	return c.Status(message.Status).JSON(fiber.Map{
		"message": message.Message,
	})
}

func GetAllProductCheckout(c *fiber.Ctx) error {
	products, message := clientService.GetAllProductCheckout(c)
	if message == nil {
		return c.Status(fiber.StatusOK).JSON(products)
	}
	return c.Status(message.Status).JSON(fiber.Map{
		"message": message.Message,
	})
}


