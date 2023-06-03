package vendorRepository

import (
	"e-commerce/db"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteProduct(ctx mongo.SessionContext, productID *primitive.ObjectID, vendorID *primitive.ObjectID) (err error) {
    query := bson.M{
        "_id":        vendorID,
    }
    update := bson.M{
        "$pull": bson.M{
            "products": productID,
        },
    }
    _, err = db.Conn.Collections["vendor"].UpdateOne(ctx, query, update)
    if err != nil {
        log.Println("An error occurred during deletion")
        return err
    }
    return nil
}