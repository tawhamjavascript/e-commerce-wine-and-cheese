package productService

import (
	"e-commerce/db"
	"e-commerce/model"
	productRepositoy "e-commerce/repository/product"
	vendorRepository "e-commerce/repository/vendedor"
	"e-commerce/service/messagesHttp"
	"e-commerce/validate"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func Create(c *fiber.Ctx) *messagesHttp.MessageErro {
	var productRegister model.RegisterProduct
	if err := c.BodyParser(&productRegister); err != nil {
		return messagesHttp.GetError(err)
	}

	idParamVendor := c.Params("id")
	idVendor, err := primitive.ObjectIDFromHex(idParamVendor)
	if err != nil {
		fmt.Println(err)
		return messagesHttp.GetError(err)
	}

	err = validate.ValidateRegisterProduct(&productRegister)
	if err != nil {
		return messagesHttp.GetError(err)
	}

	productDatabase := &model.Product{
		ID:          primitive.NewObjectID(),
		Name:        &productRegister.Name,
		Description: &productRegister.Description,
		Price:       &productRegister.Price,
		Photo:       &productRegister.Photo,
		Rating:      0.0,
		Vendor:      &idVendor,
		Comments:    &[]*model.Comments{},
	}


	session, erro := db.Conn.Client.StartSession()
	if erro != nil {
		return messagesHttp.GetError(erro)
	}
	defer session.EndSession(c.Context())
	opts := options.Transaction().SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	err = mongo.WithSession(c.Context(), session, func(sessionContext mongo.SessionContext) error {
		err := session.StartTransaction(opts)
		if err != nil {
			return err
			
		}
		err = productRepositoy.Create(sessionContext, productDatabase)
		if err != nil {
			session.AbortTransaction(sessionContext)
			return err
		}
		err = vendorRepository.AddProductVendor(sessionContext, &idVendor, productDatabase)
		if err != nil {
			session.AbortTransaction(sessionContext)
			return err
		}
		err = session.CommitTransaction(sessionContext)
		if err != nil {
			session.AbortTransaction(sessionContext)
			return err
		}
		return nil

	})
	if err != nil {
		return messagesHttp.GetError(err)
	}
	return nil
	

}
