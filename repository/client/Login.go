package clientRepository

import (
	"e-commerce/db"
	"e-commerce/model"

	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Login(c *fasthttp.RequestCtx, email string) (clientSignIn *model.SignInClientDatabase, err error) {
	filter := primitive.D{{Key: "email", Value: email}}
	opts := options.FindOne().SetProjection(bson.D{
		{Key: "email", Value: 1},
		{Key: "password", Value: 1},
		{Key: "_id", Value: 1},
	})

	err = db.Conn.Collections["client"].FindOne(c, filter, opts).Decode(&clientSignIn)
	if err != nil {
		return nil, err
	}
	return
}
