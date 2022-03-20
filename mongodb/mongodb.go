package mongodb

import (
	"context"
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client     *mongo.Client
	Collection *mongo.Collection
)

type MyStruct struct {
	Title string `mapstructure:"title"`
}

func Init(url, dbName, collectionName string) {
	MongoConnect(url, dbName, collectionName)
	// GetArticles(Client, Collection)
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

func GetDocuments(client *mongo.Client, collection *mongo.Collection, findOptions *options.FindOptions) []MyStruct {
	cursor, _ := collection.Find(context.TODO(), bson.D{}, findOptions)

	var articles []MyStruct
	for cursor.Next(context.TODO()) {
		var elem bson.M
		cursor.Decode(&elem)
		result := &MyStruct{}
		if err := mapstructure.Decode(elem, &result); err != nil {
			fmt.Println(err)
		}
		articles = append(articles, *result)
	}

	MongoDisconnect(client)

	return articles
}

func MongoDisconnect(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("disconnected...")
}
