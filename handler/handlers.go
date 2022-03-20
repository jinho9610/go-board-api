package handler

import (
	"encoding/json"
	"fmt"
	"go-board/mongodb"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("hello")
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	findOptions := options.Find()            // mongodb document search option
	findOptions.SetSort(bson.D{{"_id", -1}}) // set sort policy of searched result // 최신순으로
	findOptions.SetLimit(20)                 // set max number of searched documents policy // 최대 20개 조회

	articles := mongodb.GetDocuments(mongodb.Client, mongodb.Collection, findOptions)

	for _, article := range articles {
		fmt.Println(article.Title)
	}

	json.NewEncoder(w).Encode(articles)
}
