package clientService

import (
	"e-commerce/db"
	"e-commerce/model"
	clientRepository "e-commerce/repository/client"
	vendorRepository "e-commerce/repository/vendedor"
	"e-commerce/service/messagesHttp"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)


func Checkout(c *fiber.Ctx) *messagesHttp.MessageErro {
	var products []model.ProductView
	if err := c.BodyParser(&products); err != nil {
		log.Println("erro do body parser:" + err.Error())
		return messagesHttp.GetError(err)
	}
	idClient, _ := primitive.ObjectIDFromHex(c.Locals("id").(string))
	session, erro := db.Conn.Client.StartSession()
	if erro != nil {
		fmt.Println("erro na criação da sessão:" + erro.Error())
		return messagesHttp.GetError(erro)
	}
	defer session.EndSession(c.Context())
	opts := options.Transaction().SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	err := mongo.WithSession(c.Context(), session, func(sessionContext mongo.SessionContext) error {	
		err := session.StartTransaction(opts)
		if err != nil {
			fmt.Println("erro na no início da transação:" + erro.Error())
			return err
			
		}
		err = clientRepository.Checkout(sessionContext, &idClient, &products)
		if err != nil {
			log.Println("erro em adicionar produtos na finalização:" + err.Error())
			session.AbortTransaction(sessionContext)
			return err
		}
		for _, product := range products {
			err = vendorRepository.AddProductVendorSold(sessionContext, &product)
			if err != nil {
				log.Println("erro em adicionar vendido ao vendedor:" + err.Error())
				break
			}
		}
		if err != nil {
			session.AbortTransaction(sessionContext)
			return err
		}
		err = clientRepository.DeleteAllProductShoppingCart(sessionContext, &idClient)
		if err != nil {
			log.Println("erro ao deletar produtos do carrinho: " + err.Error())
			session.AbortTransaction(sessionContext)
			return err
		}
		err = session.CommitTransaction(sessionContext)
		if err != nil {
			fmt.Println(erro.Error())
			session.AbortTransaction(sessionContext)
			return err
		}
		return nil

	})
	if err != nil {
		fmt.Println(err)
		return messagesHttp.GetError(err)
	}
	return nil
}
