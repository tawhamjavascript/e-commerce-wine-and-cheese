package productRepository

import (
	"e-commerce/db"
	"e-commerce/model"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func UpdateProduct(ctx *fasthttp.RequestCtx, id *primitive.ObjectID, product *model.RegisterUpdateProduct)(err error) {
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
					Key: "chapter",
					Value: product.Chapter,
				},
				{
					Key: "price",
					Value: product.Price,
				},
				{
					Key: "description",
					Value: product.Description,
				},
				{
					Key: "author",
					Value: product.Author,
				},
				{
					Key: "genre",
					Value: product.Genre,
				},
				{
					Key: "publication",
					Value: product.Publication,
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