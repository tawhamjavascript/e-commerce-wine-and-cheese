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

type MongoInstance struct {
	Client      *mongo.Client
	Db          *mongo.Database
	Collections map[string]*mongo.Collection
}

var Conn *MongoInstance

func Connect(config *config.Config) {
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s", config.USER, config.PASSWORD, config.URL)))
	if err != nil {
		panic(fmt.Sprintf("Don't create a client %v", err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		panic(fmt.Sprintf("Error to connect with database %v", err.Error()))
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(fmt.Sprintf("Connect with database did not happen %v", err.Error()))
	}

	db := client.Database("E-commerce")
	collection := make(map[string]*mongo.Collection)
	collection["vendor"] = db.Collection("Vendor")
	collection["products"] = db.Collection("Products")
	collection["client"] = db.Collection("Client")
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1}, // 1 para ascendente, -1 para descendente
		Options: options.Index().SetUnique(true),
	}

	_, err = collection["vendor"].Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		panic(fmt.Sprintf("error to created an index"))

	}
	_, err = collection["client"].Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		panic(fmt.Sprintf("error to created an index"))

	}

	Conn = &MongoInstance{
		Client:      client,
		Db:          db,
		Collections: collection,
	}
}
