package vendorRepository

import (
	"e-commerce/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddProductVendor(ctx mongo.SessionContext, id *primitive.ObjectID, idProduct *primitive.ObjectID) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$push": bson.M{"products": idProduct}}
	_, err = db.Conn.Collections["vendor"].UpdateOne(ctx, filter, update)
	return err
}
