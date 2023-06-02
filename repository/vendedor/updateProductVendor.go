package vendorRepository

import (
	"e-commerce/db"
	"e-commerce/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateProduct(ctx mongo.SessionContext, vendorID *primitive.ObjectID, product *model.Product) (err error) {
    query := bson.M{
        "_id":        vendorID,
        "products._id": product.ID,
    }
    update := bson.M{
        "$set": bson.M{
            "products.$": product,
        },
    }
    _, err = db.Conn.Collections["vendor"].UpdateOne(ctx, query, update)
    if err != nil {
        log.Println("An error occurred during update")
        return err
    }
    return nil
}