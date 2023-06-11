package clientService

import (
	"e-commerce/model"
	clientRepository "e-commerce/repository/client"
	"e-commerce/service/messagesHttp"
	"e-commerce/validate"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)


func Login(c *fiber.Ctx) *messagesHttp.MessageErro {
	var SignInClient model.SignInClient

	if err := c.BodyParser(&SignInClient); err != nil {
		return messagesHttp.GetError(err)
	}
	if err := validate.ValidateLoginClient(&SignInClient); err != nil {
		return messagesHttp.GetError(err)
	}

	client, err := clientRepository.Login(c.Context(), SignInClient.Email)
	if err != nil {
		return messagesHttp.GetError(err)
	}
	err = bcrypt.CompareHashAndPassword(client.Password, []byte(SignInClient.Password))
	if err != nil {

		return messagesHttp.GetError(err)
	}
	c.Locals("id", client.ID)
	return nil
}
