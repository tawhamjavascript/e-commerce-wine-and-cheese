package productRepository

import (
	"e-commerce/db"

	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAll(ctx *fasthttp.RequestCtx) (*mongo.Cursor, error){
	filter := bson.M{}
	cursor, err := db.Conn.Collections["products"].Find(ctx, filter)
	return cursor, err

}