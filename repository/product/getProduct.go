package productRepository

import (
	"e-commerce/db"
	"e-commerce/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func GetProduct(ctx mongo.SessionContext, id *primitive.ObjectID) (*model.ProductView, error) {
	query := bson.M{
		"_id": id,
	}
	var product model.Product
	err := db.Conn.Collections["products"].FindOne(ctx, query).Decode(&product)
	return &product, err
}