package controller

import (
	productService "e-commerce/service/product"
	vendedorService "e-commerce/service/vendedor"

	"github.com/gofiber/fiber/v2"
)

func AddVendor(c *fiber.Ctx) error {
	message := vendedorService.Create(c)
	if message == nil {
		return c.Status(fiber.StatusCreated).SendStatus(fiber.StatusCreated)
	}
	return c.Status(message.Status).SendString(message.Message)
}

func LoginVendor(c *fiber.Ctx) error {
	message := vendedorService.Login(c)
	if message != nil {
		return c.Status(message.Status).SendString(message.Message)
	}
	return c.Next()
}

func AddProductVendor(c *fiber.Ctx) error {
	message := productService.Create(c)
	if message == nil {
		return c.Status(fiber.StatusOK).SendStatus(fiber.StatusOK)
	}
	return c.Status(message.Status).SendString(message.Message)
}

func DeleteVendorProduct(c *fiber.Ctx) error {
	message := productService.Delete(c)
	if message == nil {
		return c.Status(fiber.StatusNoContent).SendStatus(fiber.StatusNoContent)
	}
	return c.Status(message.Status).SendString(message.Message)
}

func UpdateVendorProduct(c *fiber.Ctx) error {
	message := productService.UpdateProduct(c)
	if message == nil {
		return c.Status(fiber.StatusOK).SendStatus(fiber.StatusOK)
	}
	return c.Status(message.Status).SendString(message.Message)
}

func GetAllVendorProduct(c *fiber.Ctx) error {
	products, message := productService.GetAll(c)
	if message == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"products": products,
		})
	}
	return c.Status(message.Status).SendString(message.Message)
}