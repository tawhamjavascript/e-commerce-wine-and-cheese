package vendorRepository

import (
	"e-commerce/db"
	"e-commerce/model"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
)

func Update(context *fasthttp.RequestCtx, uuid *uuid.UUID, vendor *model.Vendor) (err error) {
	query := bson.D{{
		Key:   "uuid",
		Value: uuid,
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
