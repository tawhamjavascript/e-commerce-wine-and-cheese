package productRepository

import (
	"e-commerce/db"
	"e-commerce/model"

	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetAllProductVendor(ctx *fasthttp.RequestCtx, idVendor *primitive.ObjectID) (*[]*model.ProductView, error) {
	var products []*model.ProductView
	filter := bson.M{"vendor": idVendor}
	cursor, err := db.Conn.Collections["products"].Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var product model.ProductView
		err := cursor.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)

	}
	
	return &products, nil

	

	
}