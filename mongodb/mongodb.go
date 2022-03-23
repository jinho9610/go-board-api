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

type Article struct {
	Title  string `json:"title" mapstructure:"title"`
	Body   string `json:"body" mapstructure:"body"`
	Writer string `json:"writer" mapstructure:"writer"`
	Passwd string `json:"passwd" mapstructure:"passwd"`
}

func Init(url, dbName, collectionName string) {
	MongoConnect(url, dbName, collectionName)
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

func GetDocuments(client *mongo.Client, collection *mongo.Collection, findOptions *options.FindOptions) []Article {
	cursor, _ := collection.Find(context.TODO(), bson.D{}, findOptions)

	var articles []Article
	for cursor.Next(context.TODO()) {
		var elem bson.M
		cursor.Decode(&elem)
		result := &Article{}
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
