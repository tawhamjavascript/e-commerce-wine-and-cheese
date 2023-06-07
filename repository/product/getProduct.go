package productRepository

import (
	"e-commerce/db"

	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func GetProduct(ctx *fasthttp.RequestCtx, id *primitive.ObjectID) (*mongo.SingleResult) {
	query := bson.M{
		"_id": id,
	}
	return db.Conn.Collections["products"].FindOne(ctx, query)

}