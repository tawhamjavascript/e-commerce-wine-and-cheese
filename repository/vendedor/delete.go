package vendorRepository

import (
	"e-commerce/db"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func Delete(context *fasthttp.RequestCtx, id uuid.UUID) (err error) {
	query := bson.D{{Key: "uuid", Value: id}}
	_, err = db.Conn.Collections["vendor"].DeleteOne(context, &query)
	if err != nil {
		log.Println("Error deleting vendor")
	}
	return err
}
