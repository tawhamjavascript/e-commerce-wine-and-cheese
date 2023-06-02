package productRepositoy

import (
	"e-commerce/db"
	"e-commerce/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func UpdateProduct(ctx mongo.SessionContext, id *primitive.ObjectID, product *model.UpdateProduct)(err error) {
	query := bson.M{
		"_id": id,
	}
	update := bson.D {
		{
			Key: "$set",
			Value: bson.D {
				{
					Key: "name",
					Value: product.Name,
				},
				{
					Key: "description",
					Value: product.Description,
				},
				{
					Key: "price",
					Value: product.Price,
				},
				{
					Key: "photo",
					Value: product.Photo,

				},
			},
		},
	}
	_, err = db.Conn.Collections["products"].UpdateOne(ctx, query, update)
	return err
}