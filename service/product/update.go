package productService

import (
	"e-commerce/db"
	"e-commerce/model"
	productRepositoy "e-commerce/repository/product"
	vendorRepository "e-commerce/repository/vendedor"
	"e-commerce/service/messagesHttp"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func UpdateProduct(c *fiber.Ctx) *messagesHttp.MessageErro{
	var productUpdate model.UpdateProduct
	if err := c.BodyParser(&productUpdate); err != nil {
		return messagesHttp.GetError(err)
	}
	idVendor, err := primitive.ObjectIDFromHex(c.Params("idVendor"))
	if err != nil {
		return messagesHttp.GetError(err)
	}
	idProduct, err := primitive.ObjectIDFromHex(c.Params("idProduct"))
	if err != nil {
		return messagesHttp.GetError(err)
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
		err = productRepositoy.UpdateProduct(sessionContext, &idProduct, &productUpdate)
		if err != nil {
			session.AbortTransaction(sessionContext)
			return err
		}
		productDatabase, err := productRepositoy.GetProduct(sessionContext, &idProduct)
		if err != nil { 
			session.AbortTransaction(sessionContext)
			return err
		}
		err = vendorRepository.UpdateProduct(sessionContext, &idVendor, productDatabase)
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