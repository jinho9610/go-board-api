package handler

import (
	"encoding/json"
	"fmt"
	"go-board/mongodb"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

func GetArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
}

func PostArticle(w http.ResponseWriter, r *http.Request) {
	// article := mongodb.Article{}
	article := mongodb.Article{}
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(article.Title)
	fmt.Println(article.Body)
	fmt.Println(article.Writer)
	fmt.Println(article.Passwd)
}
