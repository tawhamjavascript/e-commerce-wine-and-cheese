package clientRepository

import (
	"e-commerce/db"
	"e-commerce/model"

	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddShoppingCart(c *fasthttp.RequestCtx, product *model.ProductView, id *primitive.ObjectID) (error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$push": bson.M{"shoppingCart": product}}
	_, err := db.Conn.Collections["client"].UpdateOne(c, filter, update)
	return err
}