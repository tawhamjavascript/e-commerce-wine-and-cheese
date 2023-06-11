package clientRepository

import (
	"e-commerce/db"
	"e-commerce/model"

	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllProductShoppingCart(c *fasthttp.RequestCtx, id *primitive.ObjectID) (*[]model.ProductView, error){
	var client model.Client
	filter := primitive.M{"_id": id}
	opts := options.FindOne().SetProjection(bson.D{
		{Key: "shoppingCart", Value: 1},
	})
	err := db.Conn.Collections["client"].FindOne(c, filter, opts).Decode(&client)
	if err != nil {
		return nil, err
	}
	return &client.ShoppingCart, nil
	
}