package controller

import (
	productService "e-commerce/service/product"
	vendedorService "e-commerce/service/vendedor"

	"github.com/gofiber/fiber/v2"
)

func AddVendor(c *fiber.Ctx) error {
	message := vendedorService.Create(c)
	if message == nil {
		return c.Status(201).SendStatus(201)
	}
	return c.Status(message.Status).SendString(message.Message)
}

func LoginVendor(c *fiber.Ctx) error {
	message := vendedorService.Login(c)
	if message == nil {
		return c.Status(200).SendStatus(200)
	}
	return c.Status(message.Status).SendString(message.Message)
}

func AddProductVendor(c *fiber.Ctx) error {
	message := productService.Create(c)
	if message == nil {
		return c.Status(201).SendStatus(201)
	}
	return c.Status(message.Status).SendString(message.Message)
}
