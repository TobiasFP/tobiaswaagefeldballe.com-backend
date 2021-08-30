package conn

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoDB() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI("mongodb://lttr:What4reUL00k1ng4@localhost:27017/?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&ssl=false")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("lttr")
	return db, err
}
