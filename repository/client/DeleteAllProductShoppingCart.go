package clientRepository

import (
	"e-commerce/db"
	"e-commerce/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func DeleteAllProductShoppingCart(ctx mongo.SessionContext, idClient *primitive.ObjectID) error {
	filter := bson.M{"_id": idClient}
	update := bson.M{"$set": bson.M{"shoppingCart": []model.ProductView{}}}
	_, err := db.Conn.Collections["client"].UpdateOne(ctx, filter, update)
	return err
	//TODO
}