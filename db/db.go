package db

import (
	"context"
	"e-commerce/config"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

/*
# The MongoInstance struct is used to hold a connection to a MongoDB database.
# The Client field is used to hold a connection to the MongoDB server.
# The Db field is used to hold a connection to the MongoDB database.
# The Collections field is used to hold a map of all the collections in the database.
*/

type MongoInstance struct {
	Client      *mongo.Client
	Db          *mongo.Database
	Collections map[string]*mongo.Collection
}

var Conn *MongoInstance

func Connect(config *config.Config) {
	// Create a new client for connecting to the MongoDB server
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s", config.USER, config.PASSWORD, config.URL)))
	if err != nil {
		panic(fmt.Sprintf("Don't create a client %v", err))
	}

	// Creates a new context and cancels it after 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()


	// Connect to the database
	err = client.Connect(ctx)
	if err != nil {
		panic(fmt.Sprintf("Error to connect with database %v", err.Error()))
	}


	// Ping is a function that checks if the connection with the database is possible

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(fmt.Sprintf("Connect with database did not happen %v", err.Error()))
	}

	db := client.Database("E-commerce")
	// Initialize the collection map
	collection := make(map[string]*mongo.Collection)

	// Add the "vendor" collection to the map
	collection["vendor"] = db.Collection("Vendor")

	// Add the "products" collection to the map
	collection["products"] = db.Collection("Products")

	// Add the "client" collection to the map
	collection["client"] = db.Collection("Client")


	// IndexModel defines the keys for an index, the index options and the collation used for the index.
	// IndexModel is used as a wrapper for the bson.D type to represent a BSON document in a compact way.
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1}, // 1 para ascendente, -1 para descendente
		Options: options.Index().SetUnique(true),
	}

	// Create the indexes
	_, err = collection["vendor"].Indexes().CreateOne(ctx, indexModel)
		if err != nil {
			panic(fmt.Sprintf("error to created an index"))

	}
	
	_, err = collection["client"].Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		panic(fmt.Sprintf("error to created an index"))

	}

	// This code initializes the MongoDB connection and returns a pointer to the MongoDB instance.
	Conn = &MongoInstance{
			Client:      client,
			Db:          db,
			Collections: collection,
		}
}
