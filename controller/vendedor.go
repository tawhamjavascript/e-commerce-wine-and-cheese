package controller

import (
	productService "e-commerce/service/product"
	vendedorService "e-commerce/service/vendedor"

	"log"

	"github.com/gofiber/fiber/v2"
)

func AddVendor(c *fiber.Ctx) error {
	message := vendedorService.Create(c)
	if message == nil {
		return c.Status(fiber.StatusCreated).JSON(
			fiber.Map{
				"message": "Vendor created",
			},
		)
	}
	return c.Status(message.Status).JSON(fiber.Map{
		"message": message.Message,
	})
}

func LoginVendor(c *fiber.Ctx) error {
	message := vendedorService.Login(c)
	if message != nil {
		return c.Status(message.Status).JSON(fiber.Map{
			"message": message.Message,
		})
	}
	return c.Next()
}

func AddProductVendor(c *fiber.Ctx) error {
	message := productService.Create(c)
	if message == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Product created",
	})
	}
	return c.Status(message.Status).JSON(fiber.Map{
		"message": message.Message,
	})
}

func DeleteVendorProduct(c *fiber.Ctx) error {
	message := productService.Delete(c)
	if message == nil {
		return c.Status(fiber.StatusNoContent).JSON(
			fiber.Map{
				"message": "Product deleted",
		})
	}
	return c.Status(message.Status).JSON(fiber.Map{
		"message": message.Message,
	})
}

func UpdateVendorProduct(c *fiber.Ctx) error {
	message := productService.UpdateProduct(c)
	if message == nil {
		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"message": "Product updated",
		})
	}
	return c.Status(message.Status).JSON(fiber.Map{
		"message": message.Message,
	})
}

func GetAllVendorProduct(c *fiber.Ctx) error {
	products, message := productService.GetAllProductVendor(c)
	if message == nil {
		return c.Status(fiber.StatusOK).JSON(products)
	}
	return c.Status(message.Status).JSON(fiber.Map{
		"message": message.Message,
	})
}

func GetVendorProduct(c *fiber.Ctx) error {
	product, message := productService.GetProduct(c)
	if message == nil {
		return c.Status(fiber.StatusOK).JSON(product)
	}
	return c.Status(message.Status).JSON(fiber.Map{
		"message": message.Message,
	})
}

func GetAllVendorProductSold(c *fiber.Ctx) error {
	log.Println("teste")
	products, message := productService.GetAllProductSoldVendor(c)
	log.Println(products)
	if message == nil {
		
		return c.Status(fiber.StatusOK).JSON(products)
	}
	return c.Status(message.Status).JSON(fiber.Map{
		"message": message.Message,
	})
}