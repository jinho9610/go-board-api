package mongodb

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

func Init(url, dbName, collectionName string) {
	MongoConnect(url, dbName, collectionName)
	GetArticles(Client, Collection)
}

func MongoConnect(url, dbName, collectionName string) {
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	Client = client

	fmt.Println("connected...")

	Collection = client.Database(dbName).Collection(collectionName)
}

func GetArticles(client *mongo.Client, collection *mongo.Collection) {
	cursor, _ := collection.Find(context.TODO(), bson.D{{}})

	for cursor.Next(context.TODO()) {
		var elem bson.M
		cursor.Decode(&elem)
		fmt.Println(elem)
	}

	MongoDisconnect(client)
}

func MongoDisconnect(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("disconnected...")
}
