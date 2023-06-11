package clientRepository

import (
	"e-commerce/db"
	"e-commerce/model"
	"fmt"

	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func GetAllProductCheckout(ctx *fasthttp.RequestCtx, id *primitive.ObjectID) (*[]model.ProductView, error) {
	var client model.Client
	filter := bson.M{"_id": id}
	opts := options.FindOne().SetProjection(bson.D{
		{Key: "purchased", Value: 1},
	})
	err := db.Conn.Collections["client"].FindOne(ctx, filter, opts).Decode(&client)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	
	return &client.Purchased, nil
}