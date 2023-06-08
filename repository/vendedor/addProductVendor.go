package vendorRepository

import (
	"e-commerce/db"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddProductVendor(ctx mongo.SessionContext, id *primitive.ObjectID, idProduct *primitive.ObjectID) (err error) {
	log.Println("adicionando produto ao vendedor")
	filter := bson.M{"_id": id}
	update := bson.M{"$push": bson.M{"products": idProduct}}
	_, err = db.Conn.Collections["vendor"].UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("erro ao adicionar produto ao vendedor:" + err.Error())
		return err
	}
	return nil
}
