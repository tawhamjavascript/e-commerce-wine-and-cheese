package vendorRepository

import (
	"e-commerce/db"
	"e-commerce/model"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Update(context *fasthttp.RequestCtx, id *primitive.ObjectID, vendor *model.Vendor) (err error) {
	query := bson.D{{
		Key:   "_id",
		Value: id,
	}}
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{
					Key:   "email",
					Value: vendor.Email,
				},
				{
					Key:   "nome",
					Value: vendor.Name,
				},
			},
		},
	}
	_, err = db.Conn.Collections["vendor"].UpdateOne(context, &query, &update)
	return err
}
