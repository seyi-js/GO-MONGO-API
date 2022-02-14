package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func COnnectToDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	//Ping database
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connection successful.")

	return client

}

//Client instance

var DB *mongo.Client = COnnectToDB()

//Getting DB collections

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("go-mongo").Collection(collectionName)

	return collection
}
