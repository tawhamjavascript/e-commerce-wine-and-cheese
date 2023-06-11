package clientRepository

import (
	"e-commerce/db"

	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func RemoveItemShoppingCart(c *fasthttp.RequestCtx, idProduct *primitive.ObjectID, idClient *primitive.ObjectID) (error){
	filter := bson.M{"_id": idClient}
	update := bson.M{"$pull": bson.M{"shoppingCart": bson.M{"_id": idProduct}}}
	_, err := db.Conn.Collections["client"].UpdateOne(c, filter, update)
	return err
}