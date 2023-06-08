package productService

import (
	"e-commerce/db"
	"e-commerce/model"
	productRepositoy "e-commerce/repository/product"
	vendorRepository "e-commerce/repository/vendedor"
	"e-commerce/service/messagesHttp"
	"e-commerce/validate"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func Create(c *fiber.Ctx) *messagesHttp.MessageErro {
	var productRegister model.RegisterUpdateProduct
	fmt.Println(c.Locals("id"))
	if err := c.BodyParser(&productRegister); err != nil {
		log.Println("erro ao fazer o parse do body:" + err.Error())
		return messagesHttp.GetError(err)
	}

	idVendor, _ := primitive.ObjectIDFromHex(c.Locals("id").(string))
	
	err := validate.ValidateRegisterProduct(&productRegister)
	if err != nil {
		log.Println("erro na validação dos dados:" + err.Error())
		return messagesHttp.GetError(err)
	}

	productDatabase := &model.Product{
		ID:          primitive.NewObjectID(),
		Name:        &productRegister.Name,
		Chapter:     &productRegister.Chapter,
		Price:       &productRegister.Price,
		Description: &productRegister.Description,
		Author: 	 &productRegister.Author,
		Genre: 		 &productRegister.Genre,
		Publication: &productRegister.Publication,
		Photo:       &productRegister.Photo,
		Vendor:      &idVendor,
	}


	session, erro := db.Conn.Client.StartSession()
	if erro != nil {
		fmt.Println("erro na criação da sessão:" + erro.Error())
		return messagesHttp.GetError(erro)
	}
	defer session.EndSession(c.Context())
	opts := options.Transaction().SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	err = mongo.WithSession(c.Context(), session, func(sessionContext mongo.SessionContext) error {
		err := session.StartTransaction(opts)
		if err != nil {
			fmt.Println("erro na no início da transação:" + erro.Error())
			return err
			
		}
		err = productRepositoy.Create(sessionContext, productDatabase)
		if err != nil {
			fmt.Println(erro.Error())
			session.AbortTransaction(sessionContext)
			return err
		}
		err = vendorRepository.AddProductVendor(sessionContext, &idVendor, &productDatabase.ID)
		
		if err != nil {
			
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
		fmt.Println("erro na transação:" + erro.Error())
		return messagesHttp.GetError(err)
	}
	return nil
}
