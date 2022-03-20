package mongo

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client     *mongo.Client
	Collection *mongo.Collection
)

func getMongoCollection(url string) (*mongo.Client, *mongo.Collection) {
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected...")

	collection := client.Database("board").Collection("articles")

	return client, collection
}

func getArticles(client *mongo.Client, collection *mongo.Collection) {
	cursor, _ := collection.Find(context.TODO(), bson.D{{}})

	for cursor.Next(context.TODO()) {
		var elem bson.M
		cursor.Decode(&elem)
		fmt.Println(elem)
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("disconnected...")
}
