package clientService

import (
	"e-commerce/model"
	clientRepository "e-commerce/repository/client"
	"e-commerce/service/messagesHttp"
	"e-commerce/validate"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) *messagesHttp.MessageErro{
	var registerClient model.SignUpClient
	err := c.BodyParser(&registerClient)
	if err != nil {
		return messagesHttp.GetError(err)
	}
	err = validate.ValidateRegisterClient(&registerClient)
	if err != nil {
		return messagesHttp.GetError(err)
	}
	if registerClient.Password != registerClient.ConfirmPassword {
		return &messagesHttp.MessageErro{
			Status:  400,
			Message: "Password and Confirm Password not equal",
		}
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(registerClient.Password), 14)
	client := &model.Client{
		ID: 		 primitive.NewObjectID(),
		Name:        &registerClient.Name,
		Email:       &registerClient.Email,
		Password:    &hash,
		ShoppingCart: []model.ProductView{},
		Purchased:   []model.ProductView{},
	}
	err = clientRepository.Register( c.Context(), client)

	if err != nil {
		errorMensage := messagesHttp.GetError(err)
		if errorMensage == nil {
			return &messagesHttp.MessageErro{
				Status:  500,
				Message: "Error to add client",
			}
		}
		return errorMensage
	}
	c.Locals("id", client.ID)
	return nil

}