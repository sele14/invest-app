package routes

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBClient() *mongo.Client {
	// load content of env
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	
	// get mongodb endpoint
	mdbEndpoint := os.Getenv("MONGODB_URL")
	
	clientOptions := options.Client().
	ApplyURI(mdbEndpoint)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to MongoDB")

	return client
}

// client DB instance
var Client *mongo.Client = DBClient()

func connectCollection(client *mongo.Client, collectionName string) *mongo.Collection {	

	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)
	
	return collection
}