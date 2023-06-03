package productRepository

import (
	"e-commerce/db"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func Delete(ctx mongo.SessionContext, id *primitive.ObjectID) (err error) {
	_, err = db.Conn.Collections["products"].DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Println("An error occurred during deletion")
		return err
	}
	return nil
}