package vendorRepository

import (
	"e-commerce/db"
	"e-commerce/model"
	"fmt"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FilterVendor(c *fasthttp.RequestCtx, email string) (vendorSignIn *model.SignInVendoDatabase, err error) {
	filter := primitive.D{{Key: "email", Value: email}}
	opts := options.FindOne().SetProjection(bson.D{
		{"email", 1},
		{"password", 1},
		{"_id", 1},
	})
	fmt.Println(email)

	err = db.Conn.Collections["vendor"].FindOne(c, filter, opts).Decode(&vendorSignIn)
	if err != nil {
		return nil, err
	}
	return
}
