package main

import (
	"fmt"
	"go-board/router"
	"net/http"
)

func main() {
	mongoClient, articleCollection = getMongoCollection("mongodb://admin:jinho9611%40@118.67.143.181:27017/admin")
	getArticles(mongoClient, articleCollection)
	router.Init()
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", router.R)
}
