package vendedorService

import (
	"e-commerce/model"
	vendorRepository "e-commerce/repository/vendedor"
	"e-commerce/service/messagesHttp"
	"e-commerce/validate"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) *messagesHttp.MessageErro {
	var signInVendor model.SignInVendor

	if err := c.BodyParser(&signInVendor); err != nil {
		return messagesHttp.GetError(err)
	}
	if err := validate.ValidateLoginVendor(&signInVendor); err != nil {
		return messagesHttp.GetError(err)
	}

	user, err := vendorRepository.FilterVendor(c.Context(), signInVendor.Email)
	if err != nil {
		return messagesHttp.GetError(err)
	}
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(signInVendor.Password))
	if err != nil {

		return messagesHttp.GetError(err)
	}
	c.Locals("id", user.Id)
	return nil
}
