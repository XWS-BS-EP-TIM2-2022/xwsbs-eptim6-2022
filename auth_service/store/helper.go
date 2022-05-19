package store

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func CreateMongoDBConnection(uri string) *mongo.Client {
	mongoUri := "localhost:27017" //os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI("mongodb://" + mongoUri + "/?connect=direct")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}
