package vendorRepository

import (
	"e-commerce/db"
	"e-commerce/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
func AddProductVendorSold(ctx mongo.SessionContext, product *model.ProductView) error {
	filter := bson.M{"_id": product.Vendor}
	update := bson.M{"$push": bson.M{"sold": product}}
	result, err := db.Conn.Collections["vendor"].UpdateOne(ctx, filter, update)
	if result.MatchedCount == 0 {
		err = mongo.ErrNoDocuments
	}
	return err
	//TODO
}