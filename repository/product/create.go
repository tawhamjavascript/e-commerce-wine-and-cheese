package productRepository

import (
	"e-commerce/db"
	"e-commerce/model"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
)

func Create(ctx mongo.SessionContext, product *model.Product) (err error) {
	log.Println("inserindo produto no banco de dados")
	_, err = db.Conn.Collections["products"].InsertOne(ctx, product)
	if err != nil {
		log.Println("An error ocured during a insertion")
	}
	return err

}
