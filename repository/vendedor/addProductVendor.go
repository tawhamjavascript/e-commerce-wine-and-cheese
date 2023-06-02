package vendorRepository

import (
	"e-commerce/db"
	"e-commerce/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddProductVendor(ctx mongo.SessionContext, id *primitive.ObjectID, product *model.Product) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$push": bson.M{"products": product}}
	_, err = db.Conn.Collections["vendor"].UpdateOne(ctx, filter, update)
	return err
}
