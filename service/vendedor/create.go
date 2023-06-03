package vendedorService

import (
	"e-commerce/model"
	vendorRepository "e-commerce/repository/vendedor"
	"e-commerce/service/messagesHttp"
	"e-commerce/validate"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Create(c *fiber.Ctx) (messageError *messagesHttp.MessageErro) {
	var registerVendor model.RegisterVendor
	err := c.BodyParser(&registerVendor)
	if err != nil {
		return messagesHttp.GetError(err)
	}
	err = validate.ValidateRegisterVendor(&registerVendor)
	if err != nil {
		return messagesHttp.GetError(err)
	}
	if registerVendor.Password != registerVendor.ConfirmPassword {
		return &messagesHttp.MessageErro{
			Status:  400,
			Message: "Password and Confirm Password not equal",
		}
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(registerVendor.Password), 14)

	vendor := &model.Vendor{
		ID:          primitive.NewObjectID(),
		Name:        registerVendor.Name,
		Description: registerVendor.Description,
		Email:       registerVendor.Email,
		Password:    hash,
		Products:    []primitive.ObjectID{},
		Sold:        []primitive.ObjectID{},
	}
	err = vendorRepository.Create(vendor, c.Context())
	if err != nil {
		errorMensage := messagesHttp.GetError(err)
		if errorMensage == nil {
			return &messagesHttp.MessageErro{
				Status:  500,
				Message: "Error to add vendor",
			}
		}
		return errorMensage
	}
	return nil

}
