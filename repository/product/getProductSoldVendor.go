package productRepository

import (
	"e-commerce/db"
	"e-commerce/model"
	"log"

	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func GetProductSoldVendor(c *fasthttp.RequestCtx, id *primitive.ObjectID) (*[]model.Product, error){
	var vendor model.Vendor
	filter := bson.M{"_id": id}
	opts := options.FindOne().SetProjection(bson.D{
		{Key: "sold", Value: 1},
	})

	err := db.Conn.Collections["vendor"].FindOne(c, filter, opts).Decode(&vendor)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &vendor.Sold, nil

	//TODO
}