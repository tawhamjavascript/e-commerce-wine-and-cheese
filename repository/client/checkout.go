package clientRepository

import (
	"e-commerce/db"
	"e-commerce/model"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func Checkout(ctx mongo.SessionContext, idClient *primitive.ObjectID, products *[]model.ProductView) error  {
	//TODO
	filter := bson.M{"_id": idClient}
	for _, product := range *products {
		fmt.Println(product)
		update := bson.M{"$push": bson.M{"purchased": product}}
		_, err := db.Conn.Collections["client"].UpdateOne(ctx, filter, update)
		if err != nil {
			return err
		}
		
	}
	return nil
}